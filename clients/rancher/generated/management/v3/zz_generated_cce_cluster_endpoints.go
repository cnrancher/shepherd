package client

const (
	CCEClusterEndpointsType      = "cceClusterEndpoints"
	CCEClusterEndpointsFieldType = "type"
	CCEClusterEndpointsFieldURL  = "url"
)

type CCEClusterEndpoints struct {
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}
