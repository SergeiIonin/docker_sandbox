package services

import (
	"context"

	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"GoDockerSandbox/domain/yaml_helper"
	"GoDockerSandbox/infra/clients/docker_compose"
	"log"
	"strings"
)

type DockerComposeManager struct {
	composeRepo   repo.ComposeRepo
}

func NewDockerComposeManager(repo repo.ComposeRepo) *DockerComposeManager {
	return &DockerComposeManager{
		composeRepo:   repo,
	}
}

func (dcm *DockerComposeManager) BuildComposeYaml(services []model.DockerService) (yaml string) {
	//_ = dcm.sanitizeImages(services) // todo
	yaml = yaml_helper.BuildComposeYaml(services)
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

func (dcm *DockerComposeManager) SaveCompose(compose model.Compose) (id string, err error) {
	return dcm.composeRepo.Save(compose)
}

func (dcm *DockerComposeManager) UpdateCompose(id string, yaml string) (string, error) {
	composeUpd, err := yaml_helper.ParseYaml(id, yaml)
	if err != nil {
		return id, err
	}
	return dcm.composeRepo.Upsert(composeUpd)
}

func (dcm *DockerComposeManager) RunDockerCompose(id string) (err error) {
	compose, err := dcm.composeRepo.Get(id)
	if err != nil {
		log.Printf("error getting compose: %s", err.Error())
		return
	}
	composeAddress, err := docker_compose.CreateDockerComposeFile(compose)
	if err != nil {
		log.Printf("error creating docker-compose.yaml: %s", err.Error())
		return
	}

	return docker_compose.RunDockerCompose(composeAddress, compose)
}

func (dcm *DockerComposeManager) StopDockerCompose(filePath string) (err error) {
	return docker_compose.StopDockerCompose(filePath)
}

func (dcm *DockerComposeManager) GetRunningComposeServices(ctx context.Context, id string) ([]string, error) {
	containers, err := docker_compose.GetRunningContainers(ctx)
	if err != nil {
		log.Printf("error getting running containers: %s", err.Error())
		return []string{}, nil
	}
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
	return sandboxContainers, nil
}

func (dcm *DockerComposeManager) DeleteCompose(id string) error {
	return dcm.composeRepo.Delete(id)
}

func (dcm *DockerComposeManager) DeleteAllComposes() error {
	return dcm.composeRepo.DeleteAll()
}
