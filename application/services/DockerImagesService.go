package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"GoDockerSandbox/infra/clients/docker"
)

type DockerImageService struct {
	dockerClient *docker.DockerClient
	imageRepo    repo.ImageRepo
}

func NewDockerImageService(repo repo.ImageRepo) *DockerImageService {
	return &DockerImageService{
		dockerClient: docker.NewDockerClient(),
		imageRepo:    repo,
	}
}

func (dis *DockerImageService) GetImages() []string {
	return dis.dockerClient.GetImages()
}

func (dis *DockerImageService) GetImagesByName(name string) []string {
	return dis.dockerClient.GetImagesByName(name)
}

func (dis *DockerImageService) Get(id string) (model.Image, error) {
	return dis.imageRepo.Get(id)
}

func (dis *DockerImageService) GetAll() ([]model.Image, error) {
	return dis.imageRepo.GetAll()
}

func (dis *DockerImageService) Save(image model.Image) error {
	return dis.imageRepo.Save(image)
}

func (dis *DockerImageService) SaveAll(images []model.Image) error {
	return dis.imageRepo.SaveAll(images)
}

func (dis *DockerImageService) Update(image model.Image) error {
	return dis.imageRepo.Update(image)
}

func (dis *DockerImageService) UpdateAll(images []model.Image) error {
	return dis.imageRepo.UpdateAll(images)
}

func (dis *DockerImageService) Delete(id string) error {
	return dis.imageRepo.Delete(id)
}

func (dis *DockerImageService) DeleteAll() error {
	return dis.imageRepo.DeleteAll()
}
