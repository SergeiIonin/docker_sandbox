package services

import (
	"os/exec"
	"testing"
)

func TestRunContainer(t *testing.T) {
	dcs := NewDockerContainersService()
	image := "cassandra:latest"
	err := dcs.RunContainer(image)
	if err != nil {
		t.Errorf("Error running container for image %s", image)
	}
}

func TestRunContainerPlain(t *testing.T) {
	image := "cassandra:latest"
	name := "--name=" + "cassandra"
	cmd := exec.Command("docker", "run", "-d", "-p", "8080:8082", "--network=kafka_default", name, image)
	err := cmd.Run()
	if err != nil {
		t.Errorf("Error running container for image %s", image)
	}
}
