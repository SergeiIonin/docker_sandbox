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

func (dc *DockerClient) GetImages() []string {
	cmd := exec.Command("docker", "image", "ls", "--format", "{{.Repository}}:{{.Tag}}")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	images := strings.Split(string(output), "\n")
	images = images[0 : len(images)-1]
	return images
}

func (dc *DockerClient) GetImagesByName(name string) []string {
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

func (dc *DockerClient) GetRunningContainers() []string {
	return dc.getContainers([]string{})
}

func (dc *DockerClient) GetAllContainers() []string {
	args := []string{"-a"}
	return dc.getContainers(args)
}

func (dc *DockerClient) getContainers(args []string) []string {
	baseArgs := []string{"container", "ls", "--format", "{{.Names}} | {{.Image}} | {{.Status}}"}
	cmd := exec.Command("docker", append(baseArgs, args...)...)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	containers := strings.Split(string(output), "\n")
	if len(containers) == 1 && containers[0] == "" {
		return []string{}
	}
	containers = containers[0 : len(containers)-1]
	return containers
}
