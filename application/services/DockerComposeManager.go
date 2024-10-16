package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"GoDockerSandbox/domain/services"
	"GoDockerSandbox/infra/clients/docker_compose"
	"fmt"
	"log"
	"os"
	"strings"
)

type DockerComposeManager struct {
	composeRepo   repo.ComposeRepo
	yamlHelper    *services.DockerComposeYamlHelper
	composeClient *docker_compose.DockerComposeClient
}

func NewDockerComposeManager(repo repo.ComposeRepo) *DockerComposeManager {
	return &DockerComposeManager{
		composeRepo:   repo,
		yamlHelper:    services.NewDockerComposeHelper(),
		composeClient: docker_compose.NewDockerComposeClient(),
	}
}

func (dcm *DockerComposeManager) BuildComposeYaml(services []model.DockerService) (yaml string) {
	//_ = dcm.sanitizeImages(services) // todo
	yaml = dcm.yamlHelper.BuildComposeYaml(services)
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
	err, composeUpd := dcm.yamlHelper.ParseYaml(id, yaml)
	if err != nil {
		return id, err
	}
	return dcm.composeRepo.Upsert(composeUpd)
}

func (dcm *DockerComposeManager) createDockerComposeFile(compose model.Compose) (composeAddress string, err error) {
	pwd, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/docker_sandboxes/%s", pwd, compose.Id)
	if err = os.MkdirAll(filePath, 0755); err != nil {
		log.Fatalf("error creating directory: %s", err.Error())
		return
	}
	composeAddress = fmt.Sprintf("%s/docker-compose.yaml", filePath)

	yaml := compose.Yaml

	err = os.WriteFile(composeAddress, []byte(yaml), 0755)
	if err != nil {
		log.Fatalf("error creating docker-compose.yaml: %s", err.Error())
		return
	}

	return composeAddress, nil
}

func (dcm *DockerComposeManager) RunDockerCompose(id string) (err error) {
	compose, err := dcm.composeRepo.Get(id)
	if err != nil {
		log.Printf("error getting compose: %s", err.Error())
		return
	}
	composeAddress, err := dcm.createDockerComposeFile(compose)
	return dcm.composeClient.RunDockerCompose(composeAddress, compose)
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
