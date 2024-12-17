package client

const (
	CCENodePublicIPType       = "cceNodePublicIP"
	CCENodePublicIPFieldCount = "count"
	CCENodePublicIPFieldEip   = "eip"
	CCENodePublicIPFieldIDs   = "ids"
)

type CCENodePublicIP struct {
	Count int64    `json:"count,omitempty" yaml:"count,omitempty"`
	Eip   *CCEEip  `json:"eip,omitempty" yaml:"eip,omitempty"`
	IDs   []string `json:"ids,omitempty" yaml:"ids,omitempty"`
}
