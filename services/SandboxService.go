package services

import (
	"log"
)

type SandboxService struct {
	cache ImageCache
}

func NewSandboxService(cache ImageCache) *SandboxService {
	return &SandboxService{cache: cache}
}

func (ds *SandboxService) SaveAppImages(id string, images []string) error {
	err := ds.cache.SaveAll(id, images)
	if err != nil {
		log.Fatalf("SaveAppImages failed with %s\n", err)
	}
	return err
}

func (ds *SandboxService) GetAppImages(id string) ([]string, error) {
	images, err := ds.cache.GetAll(id)
	if err != nil {
		log.Fatalf("GetAppImages failed with %s\n", err)
	}
	return images, err
}

func (ds *SandboxService) UpdateAppImages(id string, images []string) error {
	err := ds.cache.UpdateAll(id, images)
	if err != nil {
		log.Fatalf("UpdateAppImages failed with %s\n", err)
	}
	return err
}

func (ds *SandboxService) DeleteAppImages(id string) error {
	err := ds.cache.DeleteAll(id)
	if err != nil {
		log.Fatalf("DeleteAppImages failed with %s\n", err)
	}
	return err	
}