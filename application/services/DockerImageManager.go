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

func (dcm *DockerImageManager) GetImages() []string {
	return dcm.dockerClient.GetImages()
}

func (dcm *DockerImageManager) GetImagesByName(name string) []string {
	return dcm.dockerClient.GetImagesByName(name)
}
