package huawei

import (
	"github.com/rancher/shepherd/clients/rancher"
	v1 "github.com/rancher/shepherd/clients/rancher/v1"
	"github.com/rancher/shepherd/extensions/cloudcredentials"
	"github.com/rancher/shepherd/extensions/defaults"
	"github.com/rancher/shepherd/extensions/defaults/namespaces"
	"github.com/rancher/shepherd/extensions/defaults/providers"
	"github.com/rancher/shepherd/extensions/defaults/stevetypes"
	"github.com/rancher/shepherd/extensions/steve"
	"github.com/rancher/shepherd/pkg/namegenerator"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateHuaweiCloudCredentials is a helper function that takes the rancher Client as a parameter and creates
// an Huawei cloud credential, and returns the CloudCredential response
func CreateHuaweiCloudCredentials(client *rancher.Client, credentials cloudcredentials.CloudCredential) (*v1.SteveAPIObject, error) {
	secretName := namegenerator.AppendRandomString(providers.Huawei)
	spec := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: cloudcredentials.GeneratedName,
			Namespace:    namespaces.CattleData,
			Annotations: map[string]string{
				"provisioning.cattle.io/driver": providers.Huawei,
				"field.cattle.io/name":          secretName,
				"field.cattle.io/creatorId":     client.UserID,
			},
		},
		Data: map[string][]byte{
			"huaweicredentialConfig-accessKey": []byte(credentials.HuaweiCredentialConfig.AccessKey),
			"huaweicredentialConfig-secretKey": []byte(credentials.HuaweiCredentialConfig.SecretKey),
			"huaweicredentialConfig-regionID":  []byte(credentials.HuaweiCredentialConfig.RegionID),
			"huaweicredentialConfig-projectID": []byte(credentials.HuaweiCredentialConfig.ProjectID),
		},
		Type: corev1.SecretTypeOpaque,
	}

	huaweiCloudCredentials, err := steve.CreateAndWaitForResource(client, stevetypes.Secret, spec, true, defaults.FiveSecondTimeout, defaults.FiveMinuteTimeout)
	if err != nil {
		return nil, err
	}

	return huaweiCloudCredentials, nil
}
