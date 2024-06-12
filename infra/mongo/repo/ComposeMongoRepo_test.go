package repo

import (
	"GoDockerSandbox/domain/model"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"testing"
	"time"
)

func TestComposeMongoRepo_Save(t *testing.T) {
	pwd, _ := os.Getwd()
	dockerComposeDir := fmt.Sprintf("%s/%s", pwd, "/docker_test/docker-compose.yaml")
	identifier := tc.StackIdentifier("compose_mongo_repo_test")
	mongoCompose, err := tc.NewDockerComposeWith(tc.WithStackFiles(dockerComposeDir), identifier)

	if err != nil {
		t.Fatal("Failed to create compose: ", err)
	}

	mongoCompose.Up(context.Background())
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

	for _, tc := range testCases {
		t.Logf("Running test case: %s\n", tc.name)
		t.Run(tc.name, func(t *testing.T) {
			id, err := repo.Save(tc.compose)
			t.Logf("Saved compose with id: %s\n", id)
			if (err != nil) != tc.wantErr {
				t.Errorf("ComposeMongoRepo.Save() error = %v, wantErr %v", err, tc.wantErr)
			}
			compose, err := repo.Get(id)
			if err != nil || compose.Name != "test_0" {
				t.Errorf("Failed to get compose: %v", err)
			}
		})
	}
}
