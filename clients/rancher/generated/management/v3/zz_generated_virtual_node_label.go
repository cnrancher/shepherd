package client

const (
	VirtualNodeLabelType       = "virtualNodeLabel"
	VirtualNodeLabelFieldName  = "name"
	VirtualNodeLabelFieldValue = "value"
)

type VirtualNodeLabel struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
