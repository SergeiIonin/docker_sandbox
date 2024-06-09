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

func (dcyh *DockerComposeYamlHelper) buildEnvironmentYaml(environment map[string]string) (envYaml string) {
	envYaml =
		indent4 + "environment:\n"
	for key, value := range environment {
		envYaml += indent6 + key + ": " + value + "\n"
	}
	return
}

func (dcyh *DockerComposeYamlHelper) buildPortsYaml(ports []string) (portsYaml string) {
	portsYaml =
		indent4 + "ports:\n"
	for _, port := range ports {
		portsYaml += indent6 + "- " + port + "\n"
	}
	return
}

// build networks yaml
func (dcyh *DockerComposeYamlHelper) buildNetworksYaml(networks []string) (networksYaml string) {
	networksYaml =
		indent4 + "networks:\n"
	for _, network := range networks {
		networksYaml += indent6 + "- " + network + "\n"
	}
	return
}

func (dcyh *DockerComposeYamlHelper) buildServiceYaml(service model.DockerService) (srvYaml string) {
	srvYaml =
		indent2 + service.Name + ":\n" +
			indent4 + "image: " + service.ImageName + "\n" +
			dcyh.buildPortsYaml(service.Ports) +
			dcyh.buildEnvironmentYaml(service.Environment) +
			dcyh.buildNetworksYaml(service.Networks)
	return
}

func (dcyh *DockerComposeYamlHelper) buildExternalNetworksYaml(services []model.DockerService) (extNetworksYaml string) {
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

func (dcyh *DockerComposeYamlHelper) BuildComposeYaml(services []model.DockerService) (composeYaml string) {
	composeYaml =
		"\n" +
			"version: '3.8'\n" +
			"\n" +
			"services:\n"

	for _, service := range services {
		composeYaml += dcyh.buildServiceYaml(service)
		composeYaml += "\n"
	}

	composeYaml +=
		dcyh.buildExternalNetworksYaml(services)

	return composeYaml
}
