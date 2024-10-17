package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/test_utils"
	"GoDockerSandbox/infra/mongo/repo"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestDockerComposeManager_Save(t *testing.T) {
	testUtils := test_utils.NewTestUtils()
	pwd, _ := os.Getwd()
	projectRoot := filepath.Dir(filepath.Dir(pwd))
	dockerComposeDir := fmt.Sprintf("%s/infra/docker_test/docker-compose.yaml", projectRoot)
	identifier := tc.StackIdentifier("compose_manager_test")
	mongoCompose, err := tc.NewDockerComposeWith(tc.WithStackFiles(dockerComposeDir), identifier)

	if err != nil {
		t.Fatal("Failed to create compose: ", err)
	}

	if err = mongoCompose.Up(context.Background()); err != nil {
		t.Fatal("Failed to start compose: ", err)
	}

	t.Cleanup(func() {
		assert.NoError(t, mongoCompose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	mongoPort := 2727
	mongoUri := fmt.Sprintf("mongodb://localhost:%d", mongoPort)

	err = mongoCompose.WaitForService("mongo_test", wait.ForListeningPort("27017").WithStartupTimeout(30*time.Second)).Up(context.Background(), tc.Wait(true))

	if err != nil {
		t.Fatal("Mongo is not accessible: ", err)
	}

	repo, err := repo.NewComposeMongoRepo(mongoUri)
	if err != nil {
		t.Fatalf("Failed to initialize ComposeMongoRepo: %v", err)
	}

	composeManager := NewDockerComposeManager(repo)

	testYaml := `
version: '3.8'

services:
  test_service_0:
    image: test_image_0
    ports:
      - 8080:8080
    environment:
      ENV_0: VALUE_0
    networks:
      - test_network_0

  test_service_1:
    image: test_image_1
    ports:
      - 9090:9090
      - 1234:1290
    environment:
      ENV_1: VALUE_1
      ENV_2: VALUE_2
    networks:
      - test_network_0
      - test_network_1

networks:
  test_network_0:
    external: true
  test_network_1:
    external: true
`

	composeTest :=
		model.Compose{
			Id:          "my-compose",
			Name:        "my-compose",
			Services:    []string{"test_service_0", "test_service_1"},
			Networks:    nil,
			AppImages:   nil,
			InfraImages: nil,
			Yaml:        testYaml,
		}

	id, err := composeManager.SaveCompose(composeTest)

	if err != nil {
		t.Fatalf("Failed to save compose: %v", err)
	}

	testUpdYaml := `
version: '3.8'

services:
  test_service_3:
    image: test_image_3
    ports:
      - 8080:8082
      - 4545:3434
    environment:
      ENV_0: VALUE_0
    networks:
      - test_network_3
      - test_network_4

  test_service_4:
    image: test_image_4
    ports:
      - 9090:9090
    environment:
      ENV_1: VALUE_1
      ENV_2: VALUE_2
    networks:
      - test_network_3

networks:
  test_network_3:
    external: true
  test_network_4:
    external: true
`

	idUpd, err := composeManager.UpdateCompose(id, testUpdYaml)
	if err != nil {
		t.Fatalf("Failed to update compose: %v", err)
	}
	composeUpd, err := composeManager.GetCompose(idUpd)
	if err != nil {
		t.Fatalf("Failed to get compose: %v", err)
	}
	servicesExpected := []string{"test_service_3", "test_service_4"}
	testUtils.CompareSlices(composeUpd.Services, servicesExpected, true, t)
	if composeUpd.Yaml != testUpdYaml {
		t.Errorf("Compose does not match")
	}
}
