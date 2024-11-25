package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/rancher/shepherd/pkg/config"
)

// ECSClient is a struct that wraps the needed ALIYUNECSConfig object, and ecs.ECs which makes the actual calls to aliyun.
type ECSClient struct {
	Client       *ecs.Client
	ClientConfig *ALIYUNECSConfigs
}

// NewECSClient is a constructor that creates an *ECSClient which a wrapper for a "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs" session and
// the aliyun ecs config.
func NewECSClient() (*ECSClient, error) {
	ecsConfig := sdk.NewConfig()
	aliyunecsClientConfig := new(ALIYUNECSConfigs)
	config.LoadConfig(ConfigurationFileKey, aliyunecsClientConfig)
	credential := credentials.NewAccessKeyCredential(aliyunecsClientConfig.ALIYUNECSAccessKeyID, aliyunecsClientConfig.ALIYUNECSSecretAccessKey)
	client, err := ecs.NewClientWithOptions(aliyunecsClientConfig.Region, ecsConfig, credential)
	if err != nil {
		return nil, err
	}
	return &ECSClient{
		Client:       client,
		ClientConfig: aliyunecsClientConfig,
	}, nil
}
