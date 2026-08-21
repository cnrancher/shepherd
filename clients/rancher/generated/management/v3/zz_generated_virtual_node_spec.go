package client

const (
	VirtualNodeSpecType             = "virtualNodeSpec"
	VirtualNodeSpecFieldDisplayName = "displayName"
	VirtualNodeSpecFieldSubnetId    = "subnetId"
	VirtualNodeSpecFieldTags        = "tags"
)

type VirtualNodeSpec struct {
	DisplayName string           `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	SubnetId    string           `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
	Tags        []VirtualNodeTag `json:"tags,omitempty" yaml:"tags,omitempty"`
}
