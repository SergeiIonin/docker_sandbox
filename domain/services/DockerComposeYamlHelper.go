package services

import (
	"GoDockerSandbox/domain/model"
	"gopkg.in/yaml.v2"
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
	if (environment == nil) || (len(environment) == 0) {
		return
	}
	envYaml =
		indent4 + "environment:\n"
	for key, value := range environment {
		envYaml += indent6 + key + ": " + value + "\n"
	}
	return
}

func (dcyh *DockerComposeYamlHelper) buildPortsYaml(ports []string) (portsYaml string) {
	if len(ports) == 0 {
		return
	}
	portsYaml = indent4 + "ports:\n"
	for _, port := range ports {
		portsYaml += indent6 + "- " + port + "\n"
	}
	return
}

// build networks yaml
func (dcyh *DockerComposeYamlHelper) buildNetworksYaml(networks []string) (networksYaml string) {
	if len(networks) == 0 {
		return
	}
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
	if len(nets) == 0 {
		return
	}

	extNetworksYaml = "networks:\n"
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

func (dcyh *DockerComposeYamlHelper) ParseYaml(id string, composeYaml string) (err error, compose model.Compose) {
	var compRaw composeRaw
	err = yaml.Unmarshal([]byte(composeYaml), &compRaw)

	compose.Name = id
	compose.Id = id
	compose.Yaml = composeYaml

	compose.Networks = make([]string, 0, len(compRaw.Networks))
	for net, _ := range compRaw.Networks {
		compose.Networks = append(compose.Networks, net)
	}

	compose.Services = make([]string, 0, len(compRaw.Services))
	for serviceName, _ := range compRaw.Services {
		compose.Services = append(compose.Services, serviceName)
	}

	return
}

type composeRaw struct {
	Version  string
	Services map[string]serviceRaw
	Networks map[string]networkRaw
}

type serviceRaw struct {
	Image       string
	Ports       []string
	Environment map[string]string
	Networks    []string
}

type networkRaw struct {
	External bool
}
