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

func NewDockerComposeManager(repo repo.ComposeRepo) *DockerComposeManager {
	return &DockerComposeManager{
		composeRepo:   repo,
		yamlBuilder:   services.NewDockerComposeHelper(),
		composeClient: docker_compose.NewDockerComposeClient(),
	}
}

func (dcm *DockerComposeManager) BuildComposeYaml(services []model.DockerService) (yaml string) {
	//_ = dcm.sanitizeImages(services) // todo
	yaml = dcm.yamlBuilder.BuildComposeYaml(services)
	return
}

// todo
// sanitize images for the case when user just clicked on the buttons to add ports/networks/envs without entering any data

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

func (dcm *DockerComposeManager) GetRunningComposeServices(id string) []string {
	containers := dcm.composeClient.GetRunningContainers()
	for i, container := range containers {
		containers[i] = container[:strings.Index(container, " | ")]
	}
	sandboxContainers := make([]string, 0, len(containers))

	for _, container := range containers {
		if strings.HasPrefix(container, id) {
			sandboxContainers = append(sandboxContainers, container)
		}
	}

	log.Printf("sandbox %s containers: %v", id, sandboxContainers)
	return sandboxContainers
}

func (dcm *DockerComposeManager) DeleteCompose(id string) error {
	return dcm.composeRepo.Delete(id)
}

func (dcm *DockerComposeManager) DeleteAllComposes() error {
	return dcm.composeRepo.DeleteAll()
}
