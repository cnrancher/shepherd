package cluster

import (
	"context"

	"github.com/rancher/shepherd/clients/rancher"
	"github.com/rancher/wrangler/pkg/summary"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// IsClusterActive is a helper function that uses the dynamic client to return cluster's ready state.
func IsClusterActive(client *rancher.Client, clusterID string) (ready bool, err error) {
	dynamic, err := client.GetRancherDynamicClient()
	if err != nil {
		logrus.Warnf("IsClusterActive: Failed to GetRancherDynamicClient: %v", err) // PANDARIA
		return
	}

	unstructuredCluster, err := dynamic.Resource(schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "clusters"}).Get(context.TODO(), clusterID, metav1.GetOptions{})
	if err != nil {
		logrus.Warnf("IsClusterActive: Failed to get cluster: %v", err) // PANDARIA
		return
	}

	summarized := summary.Summarize(unstructuredCluster)

	return summarized.IsReady(), nil
}
