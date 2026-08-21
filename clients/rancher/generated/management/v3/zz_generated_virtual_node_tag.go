package client

const (
	VirtualNodeTagType       = "virtualNodeTag"
	VirtualNodeTagFieldKey   = "key"
	VirtualNodeTagFieldValue = "value"
)

type VirtualNodeTag struct {
	Key   string `json:"key,omitempty" yaml:"key,omitempty"`
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
