package docker_compose

import (
	"context"

	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/infra/clients/docker"
	"log"
	"os/exec"
)

type DockerComposeClient struct {
	dc *docker.DockerClient
}

func NewDockerComposeClient() *DockerComposeClient {
	dc := docker.NewDockerClient()
	return &DockerComposeClient{
		dc: dc,
	}
}

func (dcc *DockerComposeClient) createNetworks(compose model.Compose) (err error) {
	nets := compose.Networks
	for _, net := range nets {
		_, err = dcc.dc.CreateNetwork(net)
		if err != nil {
			return
		}
	}
	return nil
}

func (dcc *DockerComposeClient) RunDockerCompose(composeAddress string, compose model.Compose) (err error) {
	if err = dcc.createNetworks(compose); err != nil {
		log.Printf("Error creating networks: %v", err)
		return
	}

	log.Printf("Running docker-compose for %s", composeAddress)
	cmd := exec.Command("docker-compose", "-f", composeAddress, "up", "-d")
	err = cmd.Run()
	if err != nil {
		log.Printf("Error running docker-compose: %v", err)
		return err
	}
	return nil
}

func (dcc *DockerComposeClient) StopDockerCompose(filePath string) error {
	cmd := exec.Command("docker-compose", "-f", filePath, "down")
	err := cmd.Run()
	if err != nil {
		log.Printf("Error stopping docker-compose: %v", err)
		return err
	}
	return nil
}

func (dcc *DockerComposeClient) GetRunningContainers(ctx context.Context) ([]string, error) {
	return dcc.dc.GetRunningContainers(ctx)
}
