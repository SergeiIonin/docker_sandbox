package services

type ImageCache interface {
	GetAll(id string) ([]string, error)
	SaveAll(id string, images []string) (error)
	UpdateAll(id string, images []string) error
	DeleteAll(id string) error
}
