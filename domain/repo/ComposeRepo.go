package repo

import (
	"GoDockerSandbox/domain/model"
)

type ComposeRepo interface {
	Get(id string) (model.Compose, error)
	GetAll() ([]model.Compose, error)

	Save(compose model.Compose) (string, error)
	Upsert(compose model.Compose) (string, error)

	Update(id string, yaml string) (string, error)

	Delete(id string) error
	DeleteAll() error
}
