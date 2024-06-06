package services

import (
	"GoDockerSandbox/domain/model"
)

type DockerComposeYamlHelper struct {
}

const indent2 = "  "
const indent4 = indent2 + indent2
const indent6 = indent4 + indent2

func NewDockerComposeHelper() *DockerComposeYamlHelper {
	return &DockerComposeYamlHelper{}
}

func (dcb *DockerComposeYamlHelper) buildEnvironmentYaml(environment map[string]string) (envYaml string) {
	envYaml =
		indent4 + "environment:\n"
	for key, value := range environment {
		envYaml += indent6 + key + ":" + value + "\n"
	}
	return
}

func (dcb *DockerComposeYamlHelper) buildPortsYaml(ports []string) (portsYaml string) {
	portsYaml =
		indent4 + "ports:\n"
	for _, port := range ports {
		portsYaml += indent6 + "- " + port + "\n"
	}
	return
}

// build networks yaml
func (dcb *DockerComposeYamlHelper) buildNetworksYaml(networks []string) (networksYaml string) {
	networksYaml =
		indent4 + "networks:\n"
	for _, network := range networks {
		networksYaml += indent6 + "- " + network + "\n"
	}
	return
}

func (dcb *DockerComposeYamlHelper) buildServiceYaml(service model.DockerService) (srvYaml string) {
	srvYaml =
		indent2 + service.Name + ":\n" +
			indent4 + "image: " + service.ImageName + "\n" +
			dcb.buildPortsYaml(service.Ports) +
			dcb.buildEnvironmentYaml(service.Environment) +
			dcb.buildNetworksYaml(service.Networks)
	return
}

func (dcb *DockerComposeYamlHelper) buildExternalNetworksYaml(services []model.DockerService) (extNetworksYaml string) {
	nets := make(map[string]bool)
	for _, service := range services {
		for _, network := range service.Networks {
			nets[network] = true
		}
	}

	extNetworksYaml =
		"networks:\n"
	for network := range nets {
		extNetworksYaml += indent2 + network + ":\n" +
			indent4 + "external: true\n"
	}
	return
}

func (dcb *DockerComposeYamlHelper) BuildComposeYaml(services []model.DockerService) (composeYaml string) {
	composeYaml =
		"\n" +
			"version: '3.8'\n" +
			"\n" +
			"services:\n"

	for _, service := range services {
		composeYaml += dcb.buildServiceYaml(service)
		composeYaml += "\n"
	}

	composeYaml +=
		dcb.buildExternalNetworksYaml(services)

	return composeYaml
}
