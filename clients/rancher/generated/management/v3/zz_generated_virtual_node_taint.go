package client

const (
	VirtualNodeTaintType        = "virtualNodeTaint"
	VirtualNodeTaintFieldEffect = "effect"
	VirtualNodeTaintFieldKey    = "key"
	VirtualNodeTaintFieldValue  = "value"
)

type VirtualNodeTaint struct {
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`
	Key    string `json:"key,omitempty" yaml:"key,omitempty"`
	Value  string `json:"value,omitempty" yaml:"value,omitempty"`
}
