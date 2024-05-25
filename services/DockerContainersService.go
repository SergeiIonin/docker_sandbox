package services

import (
	"log"
	"os/exec"
	"strings"
)

type DockerContainersService struct {
}

func (ds *DockerContainersService) GetAllContainers() []string {
	cmd := exec.Command("docker", "ps", "--format", "{{.Names}}")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	containers := strings.Split(string(output), "\n")
	return containers
}

func (ds *DockerContainersService) RunContainer(image string) error {
	// concat string --name=name
	name := "--name="+"my_cont"
	cmd := exec.Command("docker", "run", "-d", name, "--network=kafka_default", image)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return err
}

// stop container by name
func (ds *DockerContainersService) StopContainer(image string) {
	cmd := exec.Command("docker", "stop", image)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func NewDockerContainersService() *DockerContainersService {
	return &DockerContainersService{}
}