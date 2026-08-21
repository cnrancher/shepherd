package client

const (
	VirtualNodePoolDetailType                    = "virtualNodePoolDetail"
	VirtualNodePoolDetailFieldDeletionProtection = "deletionProtection"
	VirtualNodePoolDetailFieldLabels             = "labels"
	VirtualNodePoolDetailFieldName               = "name"
	VirtualNodePoolDetailFieldNodePoolID         = "nodePoolId"
	VirtualNodePoolDetailFieldOS                 = "os"
	VirtualNodePoolDetailFieldSecurityGroupIDs   = "securityGroupIds"
	VirtualNodePoolDetailFieldSubnetIDs          = "subnetIds"
	VirtualNodePoolDetailFieldTaints             = "taints"
	VirtualNodePoolDetailFieldVirtualNodes       = "virtualNodes"
)

type VirtualNodePoolDetail struct {
	DeletionProtection *bool              `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`
	Labels             []VirtualNodeLabel `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name               string             `json:"name,omitempty" yaml:"name,omitempty"`
	NodePoolID         string             `json:"nodePoolId,omitempty" yaml:"nodePoolId,omitempty"`
	OS                 string             `json:"os,omitempty" yaml:"os,omitempty"`
	SecurityGroupIDs   []string           `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`
	SubnetIDs          []string           `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
	Taints             []VirtualNodeTaint `json:"taints,omitempty" yaml:"taints,omitempty"`
	VirtualNodes       []VirtualNodeSpec  `json:"virtualNodes,omitempty" yaml:"virtualNodes,omitempty"`
}
