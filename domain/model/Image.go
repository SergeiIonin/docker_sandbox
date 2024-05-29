package model

type Image struct {
	Id        string `json:"id"`
	ImageName string `json:"image_name"`
	Name      string `json:"name"`
	Tag       string `json:"tag"`
	IsInfra   bool   `json:"is_infra"`
	Ports     string `json:"ports"`
	Envs      []Env  `json:"envs"`
}
