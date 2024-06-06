package model

type Compose struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	AppImages   []string `json:"app_images"`
	InfraImages []string `json:"infra_images"`
	Yaml        string   `json:"yaml"`
}
