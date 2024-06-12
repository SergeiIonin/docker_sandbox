package docker

import (
	"context"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"log"
	"os/exec"
	"slices"
	"strings"
)

type DockerClient struct {
	apiClient *client.Client
}

func NewDockerClient() *DockerClient {
	dockerClient, err := client.NewClientWithOpts(client.WithVersion("1.45"))
	if err != nil {
		log.Printf("error creating docker client: %s", err.Error())
		panic(err)
	}
	return &DockerClient{
		apiClient: dockerClient,
	}
}

func (dc *DockerClient) GetImages() []string {
	cmd := exec.Command("docker", "image", "ls", "--format", "{{.Repository}}:{{.Tag}}")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	images := strings.Split(string(output), "\n")
	images = images[0 : len(images)-1]
	return images
}

func (dc *DockerClient) GetImagesByName(name string) []string {
	cmd := exec.Command("docker", "image", "ls", "--format", "{{.Repository}}:{{.Tag}}")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	images := strings.Split(string(output), "\n")
	var filteredImages []string
	for _, image := range images {
		if strings.Contains(image, name) {
			filteredImages = append(filteredImages, image)
		}
	}
	for _, image := range filteredImages {
		log.Println(image)
	}
	return filteredImages
}

func (dc *DockerClient) GetRunningContainers() []string {
	return dc.getContainers([]string{})
}

func (dc *DockerClient) GetAllContainers() []string {
	args := []string{"-a"}
	return dc.getContainers(args)
}

func (dc *DockerClient) getNetworks() []string {
	netsSummary, err := dc.apiClient.NetworkList(context.Background(), network.ListOptions{})
	if err != nil {
		log.Printf("error getting networks: %s", err.Error())
		return []string{}
	}
	nets := make([]string, 0, len(netsSummary))
	for _, netSum := range netsSummary {
		nets = append(nets, netSum.Name)
	}
	return nets
}

func (dc *DockerClient) CreateNetwork(networkName string) (net string, err error) {
	nets := dc.getNetworks()
	networkExist := slices.Contains(nets, networkName)
	if networkExist {
		log.Printf("network %s already exists", networkName)
		return networkName, nil
	}
	_, err = dc.apiClient.NetworkCreate(context.Background(), networkName, network.CreateOptions{Driver: "bridge", Scope: "local", Internal: false})
	if err != nil {
		log.Printf("error creating network: %s", err.Error())
		return "", err
	}
	return networkName, nil
}

func (dc *DockerClient) getContainers(args []string) []string {
	baseArgs := []string{"container", "ls", "--format", "{{.Names}} | {{.Image}} | {{.Status}}"}
	cmd := exec.Command("docker", append(baseArgs, args...)...)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	containers := strings.Split(string(output), "\n")
	if len(containers) == 1 && containers[0] == "" {
		return []string{}
	}
	containers = containers[0 : len(containers)-1]
	return containers
}
