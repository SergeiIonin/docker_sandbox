package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"bufio"
	"fmt"
	"os"
)

type SandboxManager struct {
	dim *DockerImageManager
	dcm *DockerComposeManager
}

func NewSandboxManager(composeRepo repo.ComposeRepo) *SandboxManager {
	dis := NewDockerImageManager()
	dcs := NewDockerComposeService(composeRepo)
	return &SandboxManager{
		dim: dis,
		dcm: dcs,
	}
}

func (sm *SandboxManager) GetImages() []string {
	return sm.dim.GetImages()
}

func (sm *SandboxManager) SaveSandbox(name string, images []model.DockerService) (id string, err error) {
	fmt.Printf("%v\n", images)
	services := make([]string, 0, len(images))
	appImageIds := make([]string, 0, len(images))
	infraImageIds := make([]string, 0, len(images))

	for _, image := range images {
		services = append(services, image.Name)
		if image.IsInfra {
			infraImageIds = append(infraImageIds, image.Id)
		} else {
			appImageIds = append(appImageIds, image.Id)
		}
	}
	yaml := sm.dcm.BuildComposeYaml(images)

	id, err = sm.dcm.SaveCompose(model.Compose{
		Id:          name,
		Name:        name,
		Services:    services,
		AppImages:   appImageIds,
		InfraImages: infraImageIds,
		Yaml:        yaml,
	})
	return
}

func (sm *SandboxManager) GetSandbox(id string) (compose model.Compose, err error) {
	compose, err = sm.dcm.GetCompose(id)
	return
}

func (sm *SandboxManager) DeleteSandbox(name string) (err error) {
	err = sm.dcm.DeleteCompose(name)
	return
}

func (sm *SandboxManager) UpdateSandbox(id string, yaml string) (string, error) {
	_, err := sm.dcm.UpdateCompose(id, yaml)
	return id, err
}

func (sm *SandboxManager) RunSandbox(id string, yaml string) (containers []string, err error) {
	filePath := fmt.Sprintf("~/docker_sandbox/%s/docker-compose.yaml", id)
	file, err := os.Create(filePath)
	defer file.Close()
	w := bufio.NewWriter(file)
	_, err = w.WriteString(yaml)
	if err != nil {
		return
	}

	err = sm.dcm.RunDockerCompose(filePath)
	if err != nil {
		return []string{}, err
	}

	compose, err := sm.dcm.GetCompose(id)
	if err != nil {
		return []string{}, err
	}
	composeServices := compose.Services

	containers = sm.dcm.composeClient.GetRunningComposeServices(composeServices)
	return
}

func (sm *SandboxManager) StopSandbox(id string) (err error) {
	filePath := fmt.Sprintf("/docker_sandbox/%s/docker-compose.yaml", id)
	err = sm.dcm.composeClient.StopDockerCompose(filePath)
	return
}
