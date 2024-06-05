package services

import (
	"GoDockerSandbox/domain/model"
	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

type DockerComposeYamlHelper struct {
}

const indent2 = "  "
const indent4 = indent2 + indent2
const indent6 = indent4 + indent2

func NewDockerComposeHelper() *DockerComposeYamlHelper {
	return &DockerComposeYamlHelper{}
}

func (dcb *DockerComposeYamlHelper) buildEnvironmentYaml(envs map[string]string) (envYaml string) {
	envYaml =
		indent4 + "environment:\n"
	for key, value := range envs {
		envYaml += indent6 + key + ":" + value + "\n"
	}
	return
}

func (dcb *DockerComposeYamlHelper) buildServiceYaml(serviceName string, image model.DockerService, network string) (srvYaml string) {
	srvYaml =
		indent2 + image.Name + ":\n" +
			indent4 + "image: " + image.ImageName + "\n" +
			indent4 + "ports:\n" +
			indent6 + "- " + image.Ports + "\n" +
			dcb.buildEnvironmentYaml(image.Envs) +
			indent4 + "networks:\n" +
			indent6 + "- " + network + "\n"
	return
}

func (dcb *DockerComposeYamlHelper) BuildComposeYaml(images []model.DockerService, network string) (composeYaml string) {
	composeYaml =
			"\n" +
			"version: '3.8'\n" +
			"\n" +
			"services:\n"

	for _, image := range images {
		composeYaml += dcb.buildServiceYaml("app"+uuid.New().String(), image, network)
		composeYaml += "\n"
	}

	composeYaml +=
		"networks:\n" +
			indent2 + network + ":\n" +
			indent4 + "external: true"

	return composeYaml
}

func (dcb *DockerComposeYamlHelper) ParseComposeYaml(composeYaml string) (compose model.Compose, err error) {
		var dockerCompose model.DockerCompose
		err = yaml.Unmarshal([]byte(composeYaml), &dockerCompose)
		if err != nil {
			return Compose{}, err
		}
	
		var appImages, infraImages []string
		var networkName string
	
		for serviceName, service := range dockerCompose.Services {
			if networkName == "" && len(service.Networks) > 0 {
				networkName = service.Networks[0]
			}
			if service.Image != "" {
				if isInfraImage(serviceName) {
					infraImages = append(infraImages, service.Image)
				} else {
					appImages = append(appImages, service.Image)
				}
			}
		}
	
		compose := Compose{
			Id:          "", // You can generate or assign an ID here
			Name:        "", // You can set the name here
			AppImages:   appImages,
			InfraImages: infraImages,
			Network:     networkName,
			Yaml:        composeYaml,
		}
	
		return compose, nil
	}
}
