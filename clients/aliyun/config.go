package aliyun

// The json/yaml config key for the ALIYUNECSConfig
const ConfigurationFileKey = "aliyunECSConfigs"

// ALIYUNECSConfigs is the AWS authentication configuration for accessing and launching ec2 instances
type ALIYUNECSConfigs struct {
	ALIYUNECSConfig          []ALIYUNECSConfig `json:"aliyunECSConfig" yaml:"aliyunECSConfig"`
	ALIYUNECSAccessKeyID     string            `json:"aliyunAccessKeyID" yaml:"aliyunAccessKeyID"`
	ALIYUNECSSecretAccessKey string            `json:"aliyunSecretAccessKey" yaml:"aliyunSecretAccessKey"`
	Region                   string            `json:"region" yaml:"region"`
}

// ALIYUNECSConfig is the instance-specific configuration needed to launch ec2 instances in AWS
type ALIYUNECSConfig struct {
	ImageID                 string   `json:"imageId" yaml:"imageId"`
	SSHPassword             string   `json:"sshPassword" yaml:"sshPassword"`
	SSHKeyPairName          string   `json:"sshKeypair" yaml:"sshKeypair"`
	SSHPrivateKeyPath       string   `json:"sshPrivateKeyPath" yaml:"sshPrivateKeyPath"`
	InstanceType            string   `json:"instanceType" yaml:"instanceType"`
	SecurityGroupId         string   `json:"securityGroupId" yaml:"securityGroupId"`
	SecurityGroupName       string   `json:"securityGroup" yaml:"securityGroup"`
	VpcId                   string   `json:"vpcId" yaml:"vpcId"`
	VSwitchId               string   `json:"vswitchId" yaml:"vswitchId"`
	Zone                    string   `json:"zone" yaml:"zone"`
	PrivateIPOnly           bool     `json:"privateAddressOnly" yaml:"privateAddressOnly"`
	InternetMaxBandwidthOut string   `json:"internetMaxBandwidth" yaml:"internetMaxBandwidth"`
	InternetChargeType      string   `json:"internetChargeType" yaml:"internetChargeType"`
	RouteCIDR               string   `json:"routeCidr" yaml:"routeCidr"`
	SLBID                   string   `json:"slbId" yaml:"slbId"`
	DiskSize                string   `json:"diskSize" yaml:"diskSize"`
	DiskFS                  string   `json:"diskFs" yaml:"diskFs"`
	DiskCategory            string   `json:"diskCategory" yaml:"diskCategory"`
	Description             string   `json:"description" yaml:"description"`
	APIEndpoint             string   `json:"apiEndpoint" yaml:"apiEndpoint"`
	SystemDiskCategory      string   `json:"systemDiskCategory" yaml:"systemDiskCategory"`
	SystemDiskSize          string   `json:"SystemDiskSize" yaml:"SystemDiskSize"`
	ResourceGroupId         string   `json:"resourceGroupId" yaml:"resourceGroupId"`
	InstanceChargeType      string   `json:"instanceChargeType" yaml:"instanceChargeType"`
	Period                  string   `json:"period" yaml:"period"`
	PeriodUnit              string   `json:"periodUnit" yaml:"periodUnit"`
	SpotStrategy            string   `json:"spotStrategy" yaml:"spotStrategy"`
	SpotPriceLimit          string   `json:"spotPriceLimit" yaml:"spotPriceLimit"`
	SpotDuration            string   `json:"spotDuration" yaml:"spotDuration"`
	AliyunUser              string   `json:"aliyunUser" yaml:"aliyunUser"`
	DockerVersion           string   `json:"dockerVersion" yaml:"dockerVersion"`
	System                  string   `json:"system" yaml:"system"`
	OpenPorts               []string `json:"openPort" yaml:"openPort"`
	Roles                   []string `json:"roles" yaml:"roles"`
}
