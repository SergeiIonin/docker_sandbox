package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"GoDockerSandbox/domain/validation"
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

func (sm *SandboxManager) GetImages(sandboxName string) (err error, images []string) {
	err, _ = sm.v.ValidateSandboxName(sandboxName)
	if err != nil {
		return err, []string{}
	}
	return nil, sm.dim.GetImages()
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