package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"GoDockerSandbox/domain/services"
	"GoDockerSandbox/infra/clients/docker_compose"
	"log"
	"strings"
)

type DockerComposeManager struct {
	composeRepo   repo.ComposeRepo
	yamlBuilder   *services.DockerComposeYamlHelper
	composeClient *docker_compose.DockerComposeClient
}

func NewDockerComposeService(repo repo.ComposeRepo) *DockerComposeManager {
	return &DockerComposeManager{
		composeRepo:   repo,
		yamlBuilder:   services.NewDockerComposeHelper(),
		composeClient: docker_compose.NewDockerComposeClient(),
	}
}

func (dcm *DockerComposeManager) BuildComposeYaml(services []model.DockerService) (yaml string) {
	yaml = dcm.yamlBuilder.BuildComposeYaml(services)
	return
}

func (dcm *DockerComposeManager) GetCompose(id string) (model.Compose, error) {
	return dcm.composeRepo.Get(id)
}

func (dcm *DockerComposeManager) GetAllComposes() ([]model.Compose, error) {
	return dcm.composeRepo.GetAll()
}

func (dcm *DockerComposeManager) SaveCompose(compose model.Compose) (string, error) {
	return dcm.composeRepo.Save(compose)
}

func (dcm *DockerComposeManager) UpdateCompose(id string, yaml string) (string, error) {
	return dcm.composeRepo.Update(id, yaml)
}

func (dcm *DockerComposeManager) RunDockerCompose(filePath string) error {
	return dcm.composeClient.RunDockerCompose(filePath)
}

func (dcm *DockerComposeManager) GetRunningComposeServices(composeServices []string) []string {
	containers := dcm.composeClient.GetRunningContainers()
	log.Print("running containers: ")
	log.Print(containers)
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

func (dcm *DockerComposeManager) DeleteCompose(id string) error {
	return dcm.composeRepo.Delete(id)
}

func (dcm *DockerComposeManager) DeleteAllComposes() error {
	return dcm.composeRepo.DeleteAll()
}
