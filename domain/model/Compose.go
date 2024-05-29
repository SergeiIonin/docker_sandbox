package model

type Compose struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	AppImages   []string `json:"app_images"`
	InfraImages []string `json:"infra_images"`
	Network     string   `json:"network"`
	Yaml        string   `json:"yaml"`
}
