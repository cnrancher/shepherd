package pods

import (
	"errors"
	"fmt"
	"time"

	"github.com/rancher/shepherd/clients/rancher"
	v1 "github.com/rancher/shepherd/clients/rancher/v1"
	"github.com/rancher/shepherd/extensions/defaults"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	kwait "k8s.io/apimachinery/pkg/util/wait"
)

const (
	PodResourceSteveType = "pod"
)

// StatusPods is a helper function that uses the steve client to list pods on a namespace for a specific cluster
// and return the statuses in a list of strings
func StatusPods(client *rancher.Client, clusterID string) []error {
	var podErrors []error

	err := kwait.Poll(5*time.Second, defaults.FifteenMinuteTimeout, func() (done bool, err error) {
		downstreamClient, err := client.Steve.ProxyDownstream(clusterID)
		if err != nil {
			logrus.Warnf("failed to create steve doenstream proxy: %v", err) // PANDARIA:
			return false, nil
		}
		steveClient := downstreamClient.SteveType(PodResourceSteveType)

		// emptying pod errors every time we poll so that we don't return stale errors
		podErrors = []error{}

		logrus.Infof("listing downstream cluster pods...") // PANDARIA:
		pods, err := steveClient.List(nil)
		if err != nil {
			logrus.Warnf("failed to list pod: %v", err) // PANDARIA:
			// not returning the error in this case, as it could cause a false positive if we start polling too early.
			return false, nil
		}

		for _, pod := range pods.Data {
			isReady, err := IsPodReady(&pod)
			if !isReady {
				// not returning the error in this case, as it could cause a false positive if we start polling too early.
				logrus.Infof("pod [%v/%v] not ready", pod.Namespace, pod.Name) // PANDARIA:
				return false, nil
			}

			if err != nil {
				logrus.Warnf("failed to check pod ready: %v", err) // PANDARIA:
				podErrors = append(podErrors, err)
			}
		}
		return true, nil
	})

	if err != nil {
		podErrors = append(podErrors, err)
	}

	return podErrors
}

func IsPodReady(pod *v1.SteveAPIObject) (bool, error) {
	podStatus := &corev1.PodStatus{}
	err := v1.ConvertToK8sType(pod.Status, podStatus)
	if err != nil {
		return false, err
	}

	if len(podStatus.ContainerStatuses) == 0 {
		return false, nil
	}

	phase := podStatus.Phase

	if phase == corev1.PodPending {
		return false, nil
	}

	if phase == corev1.PodFailed || phase == corev1.PodUnknown {
		var errorMessage string
		for _, containerStatus := range podStatus.ContainerStatuses {
			// Rancher deploys multiple helm-operation jobs to do the same task. If one job succeeds, the others end in a terminated status.
			if containerStatus.State.Terminated == nil {
				errorMessage += fmt.Sprintf("ERROR: %s: %s, %s\n", pod.Name, podStatus.Message, podStatus.Reason)
			}
		}

		if errorMessage != "" {
			return true, errors.New(errorMessage)
		}
	}

	// Pod is running or has succeeded
	return true, nil
}
