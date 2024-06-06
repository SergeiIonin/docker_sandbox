package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
	"fmt"
)

type SandboxManager struct {
	dis *DockerImageManager
	dcs *DockerComposeManager
}

func NewSandboxManager(composeRepo repo.ComposeRepo) *SandboxManager {
	dis := NewDockerImageManager()
	dcs := NewDockerComposeService(composeRepo)
	return &SandboxManager{
		dis: dis,
		dcs: dcs,
	}
}

func (ds *SandboxManager) GetImages() []string {
	return ds.dis.GetImages()
}

func (ds *SandboxManager) SaveSandbox(name string, images []model.DockerService) (id string, err error) {
	fmt.Printf("%v\n", images)
	appImageIds := make([]string, 0, len(images))
	infraImageIds := make([]string, 0, len(images))
	for _, image := range images {
		if image.IsInfra {
			infraImageIds = append(infraImageIds, image.Id)
		} else {
			appImageIds = append(appImageIds, image.Id)
		}
	}
	yaml := ds.dcs.BuildComposeYaml(images)
	id, err = ds.dcs.composeRepo.Save(model.Compose{
		Id:          name,
		Name:        name,
		AppImages:   appImageIds,
		InfraImages: infraImageIds,
		Yaml:        yaml,
	})
	return
}

func (ds *SandboxManager) GetSandbox(name string) (compose model.Compose, err error) {
	compose, err = ds.dcs.composeRepo.Get(name)
	return
}

func (ds *SandboxManager) DeleteSandbox(name string) (err error) {
	err = ds.dcs.composeRepo.Delete(name)
	return
}

func (ds *SandboxManager) UpdateSandbox(id string, yaml string) (string, error) {
	_, err := ds.dcs.composeRepo.Update(id, yaml)
	return id, err
}

func (ds *SandboxManager) LaunchSandbox() (err error) {
	err = nil
	return
}
