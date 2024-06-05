package repo

import (
	"GoDockerSandbox/domain/model"
)

type ImageRepo interface {
	Get(id string) (model.DockerService, error)
	GetAll() ([]model.DockerService, error)

	Save(image model.DockerService) error
	SaveAll(images []model.DockerService) error

	Update(image model.DockerService) error
	UpdateAll(images []model.DockerService) error

	Delete(id string) error
	DeleteAll() error
}
