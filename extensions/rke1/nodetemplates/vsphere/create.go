package nodetemplates

import (
	"dario.cat/mergo"
	"github.com/rancher/shepherd/clients/rancher"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/rancher/shepherd/extensions/cloudcredentials/vsphere"
	"github.com/rancher/shepherd/extensions/rke1/nodetemplates"
	"github.com/rancher/shepherd/pkg/config"
)

const vmwarevsphereNodeTemplateNameBase = "vmwarevsphereNodeConfig"

// CreateVSphereNodeTemplate is a helper function that takes the rancher Client as a parameter and creates
// an VSphere node template and returns the NodeTemplate response
func CreateVSphereNodeTemplate(rancherClient *rancher.Client) (*nodetemplates.NodeTemplate, error) {
	var vmwarevsphereNodeTemplateConfig nodetemplates.VmwareVsphereNodeTemplateConfig
	config.LoadConfig(nodetemplates.VmwareVsphereNodeTemplateConfigurationFileKey, &vmwarevsphereNodeTemplateConfig)

	cloudCredential, err := vsphere.CreateVsphereCloudCredentials(rancherClient)
	if err != nil {
		return nil, err
	}

	nodeTemplate := nodetemplates.NodeTemplate{
		EngineInstallURL:                "https://releases.rancher.com/install-docker/20.10.sh",
		Name:                            vmwarevsphereNodeTemplateNameBase,
		VmwareVsphereNodeTemplateConfig: &vmwarevsphereNodeTemplateConfig,
	}

	nodeTemplateConfig := &nodetemplates.NodeTemplate{
		CloudCredentialID: cloudCredential.ID,
	}

	config.LoadConfig(nodetemplates.NodeTemplateConfigurationFileKey, nodeTemplateConfig)

	err = mergo.Merge(&nodeTemplate, nodeTemplateConfig, mergo.WithOverride)
	if err != nil {
		return nil, err
	}

	resp := &nodetemplates.NodeTemplate{}
	err = rancherClient.Management.APIBaseClient.Ops.DoCreate(management.NodeTemplateType, nodeTemplate, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetVsphereDatastoreURL is a helper to get the datastoreURL from the cloud credential object as a string
func GetVsphereDatastoreURL() string {
	var vmwarevsphereNodeTemplateConfig nodetemplates.VmwareVsphereNodeTemplateConfig

	config.LoadConfig(nodetemplates.VmwareVsphereNodeTemplateConfigurationFileKey, &vmwarevsphereNodeTemplateConfig)

	return vmwarevsphereNodeTemplateConfig.DatastoreURL
}

// GetVspherePassword is a helper to get the password from the cloud credential object as a string
func GetVspherePassword() string {
	var vmwarevsphereNodeTemplateConfig nodetemplates.VmwareVsphereNodeTemplateConfig

	config.LoadConfig(nodetemplates.VmwareVsphereNodeTemplateConfigurationFileKey, &vmwarevsphereNodeTemplateConfig)

	return vmwarevsphereNodeTemplateConfig.Password
}
