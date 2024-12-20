package docker_compose

import (
	"context"
	"fmt"

	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/infra/clients/docker"
	"log"
	"os"
	"os/exec"
)

func createNetworks(compose model.Compose) (err error) {
	nets := compose.Networks
	for _, net := range nets {
		_, err = docker.CreateNetwork(net)
		if err != nil {
			return
		}
	}
	return nil
}

func CreateDockerComposeFile(compose model.Compose) (composeAddress string, err error) {
	pwd, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s/%s", model.SandboxesDir, pwd, compose.Id)
	if err = os.MkdirAll(filePath, 0755); err != nil {
		log.Fatalf("error creating directory: %s", err.Error())
		return
	}
	composeAddress = fmt.Sprintf("%s/docker-compose.yaml", filePath)

	yaml := compose.Yaml

	err = os.WriteFile(composeAddress, []byte(yaml), 0755)
	if err != nil {
		log.Fatalf("error creating docker-compose.yaml: %s", err.Error())
		return
	}

	return composeAddress, nil
}

func RunDockerCompose(composeAddress string, compose model.Compose) (err error) {
	if err = createNetworks(compose); err != nil {
		log.Printf("Error creating networks: %v", err)
		return
	}

	log.Printf("Running docker-compose for %s", composeAddress)
	cmd := exec.Command("docker-compose", "-f", composeAddress, "up", "-d")
	err = cmd.Run()
	if err != nil {
		log.Printf("Error running docker-compose: %v", err)
		return err
	}
	return nil
}

func StopDockerCompose(filePath string) error {
	cmd := exec.Command("docker-compose", "-f", filePath, "down")
	err := cmd.Run()
	if err != nil {
		log.Printf("Error stopping docker-compose: %v", err)
		return err
	}
	return nil
}

func GetRunningContainers(ctx context.Context) ([]string, error) {
	return docker.GetRunningContainers(ctx)
}
