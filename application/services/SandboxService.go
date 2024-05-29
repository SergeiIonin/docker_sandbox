package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/repo"
)

type SandboxService struct {
	dis *DockerImageService
	dcs *DockerComposeService
}

func NewSandboxService(imageRepo repo.ImageRepo, composeRepo repo.ComposeRepo) *SandboxService {
	dis := NewDockerImageService(imageRepo)
	dcs := NewDockerComposeService(composeRepo)
	return &SandboxService{
		dis: dis,
		dcs: dcs,
	}
}

func (ds *SandboxService) GetImages() []string {
	return ds.dis.GetImages()
}

func (ds *SandboxService) SaveSandbox(name string, images []model.Image, network string) (err error) {
	appImageIds := make([]string, 0, len(images))
	infraImageIds := make([]string, 0, len(images))
	ds.dis.imageRepo.SaveAll(images)
	for _, image := range images {
		if image.IsInfra {
			infraImageIds = append(infraImageIds, image.Id)
		} else {
			appImageIds = append(appImageIds, image.Id)
		}
	}
	yaml := ds.dcs.BuildComposeYaml(images, network)
	err = ds.dcs.composeRepo.Save(model.Compose{
		Id:          name,
		Name:        name,
		AppImages:   appImageIds,
		InfraImages: infraImageIds,
		Yaml:        yaml,
	})
	return
}

func (ds *SandboxService) GetSandbox(name string) (compose model.Compose, err error) {
	compose, err = ds.dcs.composeRepo.Get(name)
	return
}

func (ds *SandboxService) DeleteSandbox(name string) (err error) {
	err = ds.dcs.composeRepo.Delete(name)
	return
}

func (ds *SandboxService) UpdateSandbox(name string, images []model.Image, network string) (err error) {
	appImageIds := make([]string, 0, len(images))
	infraImageIds := make([]string, 0, len(images))
	ds.dis.imageRepo.UpdateAll(images)
	for _, image := range images {
		if image.IsInfra {
			infraImageIds = append(infraImageIds, image.Id)
		} else {
			appImageIds = append(appImageIds, image.Id)
		}
	}
	yaml := ds.dcs.BuildComposeYaml(images, network)
	err = ds.dcs.composeRepo.Update(model.Compose{
		Id:          name,
		Name:        name,
		AppImages:   appImageIds,
		InfraImages: infraImageIds,
		Yaml:        yaml,
	})
	return
}

func (ds *SandboxService) LaunchSandbox() (err error) {
	err = nil
	return
}
