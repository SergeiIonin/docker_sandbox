package services

import (
	"GoDockerSandbox/infra/clients/docker"
)

type DockerImageService struct {
	dockerClient *docker.DockerClient
}

func NewDockerImageService() *DockerImageService {
	return &DockerImageService{
		dockerClient: docker.NewDockerClient(),
	}
}

func (dis *DockerImageService) GetImages() []string {
	return dis.dockerClient.GetImages()
}

func (dis *DockerImageService) GetImagesByName(name string) []string {
	return dis.dockerClient.GetImagesByName(name)
}
