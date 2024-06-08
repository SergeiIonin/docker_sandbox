package docker_compose

import (
	"GoDockerSandbox/infra/clients/docker"
	"log"
	"os/exec"
)

type DockerComposeClient struct {
	dc *docker.DockerClient
}

func NewDockerComposeClient() *DockerComposeClient {
	return &DockerComposeClient{}
}

func (dcc *DockerComposeClient) RunDockerCompose(filePath string) error {
	cmd := exec.Command("docker-compose", "-f", filePath, "up", "-d")
	err := cmd.Run()
	if err != nil {
		log.Printf("Error running docker-compose: %v", err)
		return err
	}
	return nil
}

func (dcc *DockerComposeClient) StopDockerCompose(filePath string) error {
	cmd := exec.Command("docker-compose", "-f", filePath, "down", "-d")
	err := cmd.Run()
	if err != nil {
		log.Printf("Error stopping docker-compose: %v", err)
		return err
	}
	return nil
}

func (dcc *DockerComposeClient) GetRunningContainers() []string {
	return dcc.dc.GetRunningContainers()
}
