package docker

import (
	"context"
	"log"
	"fmt"
	"slices"
	"strings"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

var apiClient *client.Client

func CreateDockerApiClient(version string) (err error) {
	if version == "" {
		version = "1.45"
	}
	apiClient, err = client.NewClientWithOpts(client.WithVersion(version))
	if err != nil {
		log.Printf("error creating docker client: %s", err.Error())
		return err
	}
	return nil
}

func GetImages(ctx context.Context) ([]string, error) {
	summaries, err := apiClient.ImageList(ctx, image.ListOptions{})
	if err != nil {
		log.Printf("failed to fetch images. %s\n", err.Error())
		return nil, err
	}

	images := make([]string, 0, len(summaries))
	for _, summary := range summaries {
		if len(summary.RepoTags) != 0 {
			images = append(images, summary.RepoTags...)
		}
	}

	return images, nil
}

func GetImagesByName(ctx context.Context, name string) ([]string, error) {
	images, err := GetImages(ctx)
	if err != nil {
		return []string{}, err
	}

	var filteredImages []string
	for _, image := range images {
		if strings.Contains(image, name) {
			filteredImages = append(filteredImages, image)
		}
	}
	for _, image := range filteredImages {
		log.Println(image)
	}

	return filteredImages, nil
}

func GetRunningContainers(ctx context.Context) ([]string, error) {
	return getContainers(ctx, container.ListOptions{})
}

func GetAllContainers(ctx context.Context) ([]string, error) {
	return getContainers(ctx, container.ListOptions{All: true})
}

func getNetworks() ([]string, error) {
	netsSummary, err := apiClient.NetworkList(context.Background(), network.ListOptions{})
	if err != nil {
		log.Printf("error getting networks: %s", err.Error())
		return []string{}, err
	}
	nets := make([]string, 0, len(netsSummary))
	for _, netSum := range netsSummary {
		nets = append(nets, netSum.Name)
	}
	return nets, nil
}

func CreateNetwork(networkName string) (net string, err error) {
	nets, err := getNetworks()
	if err != nil {
		return "", err
	}
	networkExist := slices.Contains(nets, networkName)
	if networkExist {
		log.Printf("network %s already exists", networkName)
		return networkName, nil
	}
	_, err = apiClient.NetworkCreate(context.Background(), networkName, network.CreateOptions{Driver: "bridge", Scope: "local", Internal: false})
	if err != nil {
		log.Printf("error creating network: %s", err.Error())
		return "", err
	}
	return networkName, nil
}

func getContainers(ctx context.Context, opts container.ListOptions) ([]string, error) {
	containers, err := apiClient.ContainerList(ctx, opts)
	if err != nil {
		log.Printf("error getting containers: %s", err.Error())
		return []string{}, err
	}

	containersList := make([]string, 0, len(containers))
	for _, container := range containers {
		containersList = append(containersList, fmt.Sprintf("%s | %s | %s", container.Names[0], container.Image, container.State))
	}

	return containersList, nil
}
