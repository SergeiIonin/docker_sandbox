package model

type DockerService struct {
	Id          string            `json:"id"`
	ImageName   string            `json:"image_name" yaml:"image"`
	Name        string            `json:"name"`
	Tag         string            `json:"tag"`
	IsInfra     bool              `json:"is_infra"`
	Ports       []string          `json:"ports" yaml:"ports"`
	Environment map[string]string `json:"environment" yaml:"environment"`
	Networks    []string          `json:"networks" yaml:"networks"`
}
