package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"GoDockerSandbox/domain/services"
	"GoDockerSandbox/infra/clients/docker_compose"
	"strings"
)

type DockerComposeManager struct {
	composeRepo   repo.ComposeRepo
	yamlBuilder   *services.DockerComposeYamlHelper
	composeClient *docker_compose.DockerComposeClient
}

func NewDockerComposeService(repo repo.ComposeRepo) *DockerComposeManager {
	return &DockerComposeManager{
		composeRepo: repo,
	}
}

func (dcs *DockerComposeManager) BuildComposeYaml(services []model.DockerService) (yaml string) {
	yaml = dcs.yamlBuilder.BuildComposeYaml(services)
	return
}

func (dcs *DockerComposeManager) GetCompose(id string) (model.Compose, error) {
	return dcs.composeRepo.Get(id)
}

func (dcs *DockerComposeManager) GetAllComposes() ([]model.Compose, error) {
	return dcs.composeRepo.GetAll()
}

func (dcs *DockerComposeManager) SaveCompose(compose model.Compose) (string, error) {
	return dcs.composeRepo.Save(compose)
}

func (dcs *DockerComposeManager) UpdateCompose(id string, yaml string) (string, error) {
	return dcs.composeRepo.Update(id, yaml)
}

func (dcs *DockerComposeManager) RunDockerCompose(filePath string) error {
	return dcs.composeClient.RunDockerCompose(filePath)
}

func (dcs *DockerComposeManager) GetRunningComposeServices(composeServices []string) []string {
	containers := dcs.composeClient.GetRunningContainers()
	composeContainers := make([]string, 0, len(composeServices))
	servicesMap := make(map[string]bool)
	for _, service := range composeServices {
		servicesMap[service] = true
	}
	var containerName string
	for _, container := range containers {
		containerName = container[:strings.Index(container, " | ")]
		if servicesMap[containerName] {
			composeContainers = append(composeContainers, container)
		}
	}
	return composeContainers
}

func (dcs *DockerComposeManager) DeleteCompose(id string) error {
	return dcs.composeRepo.Delete(id)
}

func (dcs *DockerComposeManager) DeleteAllComposes() error {
	return dcs.composeRepo.DeleteAll()
}
