package repo

import (
	"GoDockerSandbox/domain/model"
)

type ComposeRepo interface {
	Get(id string) (model.Compose, error)
	GetAll() ([]model.Compose, error)

	Save(compose model.Compose) error

	Update(compose model.Compose) error

	Delete(id string) error
	DeleteAll() error
}
