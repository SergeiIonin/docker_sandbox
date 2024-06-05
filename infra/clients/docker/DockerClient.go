package docker

import (
	"log"
	"os/exec"
	"strings"
)

type DockerClient struct {
}

func NewDockerClient() *DockerClient {
	return &DockerClient{}
}

func (ds *DockerClient) GetImages() []string {
	cmd := exec.Command("docker", "image", "ls", "--format", "{{.Repository}}:{{.Tag}}")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	images := strings.Split(string(output), "\n")
	return images
}

func (ds *DockerClient) GetImagesByName(name string) []string {
	cmd := exec.Command("docker", "image", "ls", "--format", "{{.Repository}}:{{.Tag}}")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	images := strings.Split(string(output), "\n")
	var filteredImages []string
	for _, image := range images {
		if strings.Contains(image, name) {
			filteredImages = append(filteredImages, image)
		}
	}
	for _, image := range filteredImages {
		log.Println(image)
	}
	return filteredImages
}