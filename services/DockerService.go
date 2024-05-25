package services

import (
	"GoDockerSandbox/infra"
)

type DockerService struct {
	ImagesService *DockerImageService
	ContainersService *DockerContainersService
	ComposeBuilder *DockerComposeBuilder
}

func NewDockerService() *DockerService {
	dis := NewDockerImageService()
	dcs := NewDockerContainersService()
	imageCache := infra.NewImageCacheInMemImpl()
	cs := NewDockerComposeBuilder(imageCache)
	return &DockerService {
		ImagesService: dis,
		ContainersService: dcs,
		ComposeBuilder: cs,
	}
}