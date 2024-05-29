package repo

import (
	"GoDockerSandbox/domain/model"
)

type ImageRepo interface {
	Get(id string) (model.Image, error)
	GetAll() ([]model.Image, error)

	Save(image model.Image) error
	SaveAll(images []model.Image) error

	Update(image model.Image) error
	UpdateAll(images []model.Image) error

	Delete(id string) error
	DeleteAll() error
}
