package model

type DockerCompose struct {
    Version  string             	  `yaml:"version"`
    Services map[string]DockerService `yaml:"services"`
    Networks map[string]Network 	  `yaml:"networks"`
}