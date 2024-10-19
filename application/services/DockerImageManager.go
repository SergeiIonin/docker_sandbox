package services

import (
	"GoDockerSandbox/infra/clients/docker"
	"context"
)

type DockerImageManager struct {
	dockerClient *docker.DockerClient
}

func NewDockerImageManager() *DockerImageManager {
	return &DockerImageManager{
		dockerClient: docker.NewDockerClient(),
	}
}

func (dcm *DockerImageManager) GetImages(ctx context.Context) ([]string, error) {
	return dcm.dockerClient.GetImages(ctx)
}

func (dcm *DockerImageManager) GetImagesByName(ctx context.Context, name string) ([]string, error) {
	return dcm.dockerClient.GetImagesByName(ctx, name)
}
