package services

import (
	"strings"

	"github.com/google/uuid"
)

type DockerComposeBuilder struct {
	imageCache ImageCache
}

const indent2 = "  "
const indent4 = indent2 + indent2
const indent6 = indent4 + indent2

func NewDockerComposeBuilder(imageCache ImageCache) *DockerComposeBuilder {
	return &DockerComposeBuilder{imageCache: imageCache}
}

func (dcb *DockerComposeBuilder) buildService(serviceName string, image string, network string) (srvYaml string) {
	srvYaml =
		indent2 + serviceName + ":\n" +
		indent4 + "image: " + image + "\n" +
		indent4 + "networks:\n" +
		indent6 + "- " + network + "\n"
	return
}

func (dcb *DockerComposeBuilder) BuildComposeFile(images []string, network string) (composeYaml string) {
	composeYaml =
	"version: '3.8'\n" +
	"\n" +
	"services:\n"
	
	for _, image := range images {
		composeYaml += dcb.buildService("app" + uuid.New().String(), image, network)
		composeYaml += "\n"
	}

	htmlString := strings.Replace(composeYaml, "\n", "<br>", -1)

	return htmlString
}
