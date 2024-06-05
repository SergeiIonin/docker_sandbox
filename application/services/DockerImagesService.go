package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"GoDockerSandbox/infra/clients/docker"
)

type DockerServiceManager struct {
	dockerClient *docker.DockerClient
	imageRepo    repo.ImageRepo
}

func NewDockerServiceManager(repo repo.ImageRepo) *DockerServiceManager {
	return &DockerServiceManager{
		dockerClient: docker.NewDockerClient(),
		imageRepo:    repo,
	}
}

func (dis *DockerServiceManager) GetImages() []string {
	return dis.dockerClient.GetImages()
}

func (dis *DockerServiceManager) GetImagesByName(name string) []string {
	return dis.dockerClient.GetImagesByName(name)
}

func (dis *DockerServiceManager) Get(id string) (model.DockerService, error) {
	return dis.imageRepo.Get(id)
}

func (dis *DockerServiceManager) GetAll() ([]model.DockerService, error) {
	return dis.imageRepo.GetAll()
}

func (dis *DockerServiceManager) Save(image model.DockerService) error {
	return dis.imageRepo.Save(image)
}

func (dis *DockerServiceManager) SaveAll(images []model.DockerService) error {
	return dis.imageRepo.SaveAll(images)
}

func (dis *DockerServiceManager) Update(image model.DockerService) error {
	return dis.imageRepo.Update(image)
}

func (dis *DockerServiceManager) UpdateAll(images []model.DockerService) error {
	return dis.imageRepo.UpdateAll(images)
}

func (dis *DockerServiceManager) Delete(id string) error {
	return dis.imageRepo.Delete(id)
}

func (dis *DockerServiceManager) DeleteAll() error {
	return dis.imageRepo.DeleteAll()
}
