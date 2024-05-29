package services

import (
	"GoDockerSandbox/domain/model"
	"github.com/google/uuid"
	"strings"
)

type DockerComposeYamlBuilder struct {
}

const indent2 = "  "
const indent4 = indent2 + indent2
const indent6 = indent4 + indent2

func NewDockerComposeBuilder() *DockerComposeYamlBuilder {
	return &DockerComposeYamlBuilder{}
}

func (dcb *DockerComposeYamlBuilder) buildServiceYaml(serviceName string, image model.Image, network string) (srvYaml string) {
	srvYaml =
		indent2 + image.Name + ":\n" +
			indent4 + "image: " + image.ImageName + "\n" +
			indent4 + "networks:\n" +
			indent6 + "- " + network + "\n"
	return
}

func (dcb *DockerComposeYamlBuilder) BuildComposeYaml(images []model.Image, network string) (composeYaml string) {
	composeYaml =
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

	htmlString := strings.Replace(composeYaml, "\n", "<br>", -1)

	return htmlString
}
