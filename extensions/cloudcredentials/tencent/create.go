package tencent

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

// CreateTencentCloudCredentials is a helper function that takes the rancher Client as a parameter and creates
// an Tencent cloud credential, and returns the CloudCredential response
func CreateTencentCloudCredentials(client *rancher.Client, credentials cloudcredentials.CloudCredential) (*v1.SteveAPIObject, error) {
	secretName := namegenerator.AppendRandomString(providers.Tencent)
	spec := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: cloudcredentials.GeneratedName,
			Namespace:    namespaces.CattleData,
			Annotations: map[string]string{
				"provisioning.cattle.io/driver": providers.Tencent,
				"field.cattle.io/name":          secretName,
				"field.cattle.io/creatorId":     client.UserID,
			},
		},
		Data: map[string][]byte{
			"tkecredentialConfig-accessKeyId":     []byte(credentials.TencentCredentialConfig.AccessKeyID),
			"tkecredentialConfig-accessKeySecret": []byte(credentials.TencentCredentialConfig.AccessKeySecret),
		},
		Type: corev1.SecretTypeOpaque,
	}

	tkeCloudCredentials, err := steve.CreateAndWaitForResource(client, stevetypes.Secret, spec, true, defaults.FiveSecondTimeout, defaults.FiveMinuteTimeout)
	if err != nil {
		return nil, err
	}

	return tkeCloudCredentials, nil
}
