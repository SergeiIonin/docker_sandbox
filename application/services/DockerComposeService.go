package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"GoDockerSandbox/domain/services"
)

type DockerComposeManager struct {
	composeRepo  repo.ComposeRepo
	yamlBuiilder *services.DockerComposeYamlHelper
}

func NewDockerComposeService(repo repo.ComposeRepo) *DockerComposeManager {
	return &DockerComposeManager{
		composeRepo: repo,
	}
}

func (dcs *DockerComposeManager) BuildComposeYaml(images []model.DockerService, network string) (yaml string) {
	yaml = dcs.yamlBuiilder.BuildComposeYaml(images, network)
	return
}

func (dcs *DockerComposeManager) ParseComposeYaml(yaml string) (compose model.Compose) {
	compose = dcs.ParseComposeYaml(yaml)
	return
}

func (dcs *DockerComposeManager) GetCompose(id string) (model.Compose, error) {
	return dcs.composeRepo.Get(id)
}

func (dcs *DockerComposeManager) GetAllComposes() ([]model.Compose, error) {
	return dcs.composeRepo.GetAll()
}

func (dcs *DockerComposeManager) SaveCompose(compose model.Compose) error {
	return dcs.composeRepo.Save(compose)
}

func (dcs *DockerComposeManager) UpdateCompose(compose model.Compose) error {
	return dcs.composeRepo.Update(compose)
}

func (dcs *DockerComposeManager) DeleteCompose(id string) error {
	return dcs.composeRepo.Delete(id)
}

func (dcs *DockerComposeManager) DeleteAllComposes() error {
	return dcs.composeRepo.DeleteAll()
}
