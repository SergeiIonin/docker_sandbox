package repo

import (
	"GoDockerSandbox/domain/model"
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

func TestComposeMongoRepo_Save(t *testing.T) {
	pwd, _ := os.Getwd()
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(pwd)))
	dockerComposeDir := fmt.Sprintf("%s/infra/docker_test/docker-compose.yaml", projectRoot)
	identifier := tc.StackIdentifier("compose_mongo_repo_test")
	mongoCompose, err := tc.NewDockerComposeWith(tc.WithStackFiles(dockerComposeDir), identifier)

	if err != nil {
		t.Fatal("Failed to create compose: ", err)
	}

	err = mongoCompose.Up(context.Background())
	if err != nil {
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

	repo, err := NewComposeMongoRepo(mongoUri)
	if err != nil {
		t.Fatalf("Failed to initialize ComposeMongoRepo: %v", err)
	}

	// Define the test cases
	testCases := []struct {
		name    string
		compose model.Compose
		wantErr bool
	}{
		{
			name: "Valid compose",
			compose: model.Compose{
				Id:          "test_0",
				Name:        "test_0",
				Services:    []string{"nginx"},
				AppImages:   nil,
				InfraImages: []string{"nginx"},
				Yaml:        "version: '3'\nservices:\n  web:\n    image: nginx",
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}

	for _, testCase := range testCases {
		t.Logf("Running test case: %s\n", testCase.name)
		t.Run(testCase.name, func(t *testing.T) {
			id, err := repo.Save(testCase.compose)
			t.Logf("Saved compose with id: %s\n", id)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ComposeMongoRepo.Save() error = %v, wantErr %v", err, testCase.wantErr)
			}
			compose, err := repo.Get(id)
			if err != nil || compose.Name != "test_0" {
				t.Errorf("Failed to get compose: %v", err)
			}
		})
	}
}
