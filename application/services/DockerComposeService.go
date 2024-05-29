package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"GoDockerSandbox/domain/services"
)

type DockerComposeService struct {
	composeRepo  repo.ComposeRepo
	yamlBuiilder *services.DockerComposeYamlBuilder
}

func NewDockerComposeService(repo repo.ComposeRepo) *DockerComposeService {
	return &DockerComposeService{
		composeRepo: repo,
	}
}

func (dcs *DockerComposeService) BuildComposeYaml(images []model.Image, network string) (yaml string) {
	yaml = dcs.yamlBuiilder.BuildComposeYaml(images, network)
	return
}

func (dcs *DockerComposeService) GetCompose(id string) (model.Compose, error) {
	return dcs.composeRepo.Get(id)
}

func (dcs *DockerComposeService) GetAllComposes() ([]model.Compose, error) {
	return dcs.composeRepo.GetAll()
}

func (dcs *DockerComposeService) SaveCompose(compose model.Compose) error {
	return dcs.composeRepo.Save(compose)
}

func (dcs *DockerComposeService) UpdateCompose(compose model.Compose) error {
	return dcs.composeRepo.Update(compose)
}

func (dcs *DockerComposeService) DeleteCompose(id string) error {
	return dcs.composeRepo.Delete(id)
}

func (dcs *DockerComposeService) DeleteAllComposes() error {
	return dcs.composeRepo.DeleteAll()
}
