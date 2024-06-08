package model

type Compose struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Services    []string `json:"services"`
	AppImages   []string `json:"app_images"`
	InfraImages []string `json:"infra_images"`
	Yaml        string   `json:"yaml"`
}
