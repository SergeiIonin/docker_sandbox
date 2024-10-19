package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"GoDockerSandbox/domain/validation"
	"context"
	"fmt"
	"os"
)

type SandboxManager struct {
	dim *DockerImageManager
	dcm *DockerComposeManager
	v   *validation.Validations
}

func NewSandboxManager(composeRepo repo.ComposeRepo) *SandboxManager {
	return &SandboxManager{
		dim: NewDockerImageManager(),
		dcm: NewDockerComposeManager(composeRepo),
		v:   validation.NewValidations(),
	}
}

func (sm *SandboxManager) GetImages(sandboxName string) (images []string, err error) {
	err, _ = sm.v.ValidateSandboxName(sandboxName)
	if err != nil {
		return []string{}, err
	}
	images, err = sm.dim.GetImages(context.Background())
	if err != nil {
		return []string{}, err
	}

	return images, nil
}

func (sm *SandboxManager) SaveSandbox(name string, dockerServices []model.DockerService) (id string, err error) {
	fmt.Printf("%v\n", dockerServices)
	services := make([]string, 0, len(dockerServices))
	appImageIds := make([]string, 0, len(dockerServices))
	infraImageIds := make([]string, 0, len(dockerServices))
	networks := make([]string, 0, len(dockerServices))

	getNetworks := func(service model.DockerService) []string {
		return service.Networks
	}

	for _, srv := range dockerServices {
		services = append(services, srv.Name)
		networks = append(networks, getNetworks(srv)...)
		if srv.IsInfra {
			infraImageIds = append(infraImageIds, srv.Id)
		} else {
			appImageIds = append(appImageIds, srv.Id)
		}
	}
	yaml := sm.dcm.BuildComposeYaml(dockerServices)

	id, err = sm.dcm.SaveCompose(model.Compose{
		Id:          name,
		Name:        name,
		Services:    services,
		Networks:    networks,
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

func (sm *SandboxManager) RunSandbox(id string) (err error) {
	err = sm.dcm.RunDockerCompose(id)
	return
}

func (sm *SandboxManager) GetRunningSandboxServices(id string) []string {
	return sm.dcm.GetRunningComposeServices(id)
}

func (sm *SandboxManager) StopSandbox(id string) (err error) {
	pwd, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/docker_sandboxes/%s", pwd, id)
	composeAddr := fmt.Sprintf("%s/docker-compose.yaml", filePath)

	err = sm.dcm.composeClient.StopDockerCompose(composeAddr)
	return
}
