package services

import (
	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/domain/test_utils"
	"testing"
)

func TestDockerComposeYamlHelper(t *testing.T) {
	testUtils := test_utils.NewTestUtils()
	dcyh := NewDockerComposeHelper()

	service_0 := model.DockerService{
		Id:        "test_service_0",
		ImageName: "test_image_0",
		Name:      "test_service_0",
		Tag:       "latest",
		Ports:     []string{"8080:8080"},
		Environment: map[string]string{
			"ENV_0": "VALUE_0",
		},
		Networks: []string{"test_network_0"},
	}

	service_1 := model.DockerService{
		Id:        "test_service_1",
		ImageName: "test_image_1",
		Name:      "test_service_1",
		Tag:       "some_tag",
		Ports:     []string{"9090:9090", "1234:1290"},
		Environment: map[string]string{
			"ENV_1": "VALUE_1",
			"ENV_2": "VALUE_2",
		},
		Networks: []string{"test_network_0", "test_network_1"},
	}

	services := []model.DockerService{service_0, service_1}

	composeYaml := dcyh.BuildComposeYaml(services)

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

	if composeYaml != testYaml {
		t.Errorf("got:\n%s\nwant:\n%s", composeYaml, testYaml)
	}

	composeFromYaml, err := dcyh.ParseYaml("test", testYaml)
	t.Logf("Compose from yaml: %v\n", composeFromYaml)
	testUtils.CompareSlices(composeFromYaml.Services, []string{"test_service_0", "test_service_1"}, true, t)
	if err != nil {
		t.Errorf("Failed to parse yaml: %v", err)
	}

}
