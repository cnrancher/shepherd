package client

const (
	ClusterSpecBaseType                                                      = "clusterSpecBase"
	ClusterSpecBaseFieldAgentEnvVars                                         = "agentEnvVars"
	ClusterSpecBaseFieldAgentImageOverride                                   = "agentImageOverride"
	ClusterSpecBaseFieldClusterAgentDeploymentCustomization                  = "clusterAgentDeploymentCustomization"
	ClusterSpecBaseFieldClusterSecrets                                       = "clusterSecrets"
	ClusterSpecBaseFieldDefaultClusterRoleForProjectMembers                  = "defaultClusterRoleForProjectMembers"
	ClusterSpecBaseFieldDefaultPodSecurityAdmissionConfigurationTemplateName = "defaultPodSecurityAdmissionConfigurationTemplateName"
	ClusterSpecBaseFieldDesiredAgentImage                                    = "desiredAgentImage"
	ClusterSpecBaseFieldDesiredAuthImage                                     = "desiredAuthImage"
	ClusterSpecBaseFieldDockerRootDir                                        = "dockerRootDir"
	ClusterSpecBaseFieldEnableGPUManagement                                  = "enableGPUManagement"
	ClusterSpecBaseFieldEnableNetworkPolicy                                  = "enableNetworkPolicy"
	ClusterSpecBaseFieldFleetAgentDeploymentCustomization                    = "fleetAgentDeploymentCustomization"
	ClusterSpecBaseFieldFluentdLogDir                                        = "fluentdLogDir"
	ClusterSpecBaseFieldLocalClusterAuthEndpoint                             = "localClusterAuthEndpoint"
	ClusterSpecBaseFieldRancherKubernetesEngineConfig                        = "rancherKubernetesEngineConfig"
	ClusterSpecBaseFieldSystemDefaultRegistry                                = "systemDefaultRegistry"
	ClusterSpecBaseFieldWindowsPreferedCluster                               = "windowsPreferedCluster"
)

type ClusterSpecBase struct {
	AgentEnvVars                                         []EnvVar                       `json:"agentEnvVars,omitempty" yaml:"agentEnvVars,omitempty"`
	AgentImageOverride                                   string                         `json:"agentImageOverride,omitempty" yaml:"agentImageOverride,omitempty"`
	ClusterAgentDeploymentCustomization                  *AgentDeploymentCustomization  `json:"clusterAgentDeploymentCustomization,omitempty" yaml:"clusterAgentDeploymentCustomization,omitempty"`
	ClusterSecrets                                       *ClusterSecrets                `json:"clusterSecrets,omitempty" yaml:"clusterSecrets,omitempty"`
	DefaultClusterRoleForProjectMembers                  string                         `json:"defaultClusterRoleForProjectMembers,omitempty" yaml:"defaultClusterRoleForProjectMembers,omitempty"`
	DefaultPodSecurityAdmissionConfigurationTemplateName string                         `json:"defaultPodSecurityAdmissionConfigurationTemplateName,omitempty" yaml:"defaultPodSecurityAdmissionConfigurationTemplateName,omitempty"`
	DesiredAgentImage                                    string                         `json:"desiredAgentImage,omitempty" yaml:"desiredAgentImage,omitempty"`
	DesiredAuthImage                                     string                         `json:"desiredAuthImage,omitempty" yaml:"desiredAuthImage,omitempty"`
	DockerRootDir                                        string                         `json:"dockerRootDir,omitempty" yaml:"dockerRootDir,omitempty"`
	EnableGPUManagement                                  bool                           `json:"enableGPUManagement,omitempty" yaml:"enableGPUManagement,omitempty"`
	EnableNetworkPolicy                                  *bool                          `json:"enableNetworkPolicy,omitempty" yaml:"enableNetworkPolicy,omitempty"`
	FleetAgentDeploymentCustomization                    *AgentDeploymentCustomization  `json:"fleetAgentDeploymentCustomization,omitempty" yaml:"fleetAgentDeploymentCustomization,omitempty"`
	FluentdLogDir                                        string                         `json:"fluentdLogDir,omitempty" yaml:"fluentdLogDir,omitempty"`
	LocalClusterAuthEndpoint                             *LocalClusterAuthEndpoint      `json:"localClusterAuthEndpoint,omitempty" yaml:"localClusterAuthEndpoint,omitempty"`
	RancherKubernetesEngineConfig                        *RancherKubernetesEngineConfig `json:"rancherKubernetesEngineConfig,omitempty" yaml:"rancherKubernetesEngineConfig,omitempty"`
	SystemDefaultRegistry                                string                         `json:"systemDefaultRegistry,omitempty" yaml:"systemDefaultRegistry,omitempty"`
	WindowsPreferedCluster                               bool                           `json:"windowsPreferedCluster,omitempty" yaml:"windowsPreferedCluster,omitempty"`
}
