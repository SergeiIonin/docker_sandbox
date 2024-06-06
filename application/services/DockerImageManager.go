package services

import (
	"GoDockerSandbox/infra/clients/docker"
)

type DockerImageManager struct {
	dockerClient *docker.DockerClient
}

func NewDockerImageManager() *DockerImageManager {
	return &DockerImageManager{
		dockerClient: docker.NewDockerClient(),
	}
}

func (dis *DockerImageManager) GetImages() []string {
	return dis.dockerClient.GetImages()
}

func (dis *DockerImageManager) GetImagesByName(name string) []string {
	return dis.dockerClient.GetImagesByName(name)
}
