package docker

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

func TestDockerClientNew(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DockerClient Suite")
}

var _ = Describe("DockerClient", func() {
	dockerClient := NewDockerClient()

	Describe("GetImages", func() {
		It("should return a list of images", func() {
			images, err := dockerClient.GetImages(context.Background())
			for _, image := range images {
				log.Printf("image: %s\n", image)
			}
			Expect(err).NotTo(HaveOccurred())
			Expect(len(images)).To(BeNumerically(">", 0))
		})

		It("should return a list of images matching some name", func() {
			images, err := dockerClient.GetImagesByName(context.Background(), "mongo")
			for _, image := range images {
				log.Printf("image: %s\n", image)
			}
			Expect(err).NotTo(HaveOccurred())
			Expect(len(images)).To(BeNumerically(">", 0))
		})
	})

	Describe("GetContainers", func() {

		rawDockerClient, err := client.NewClientWithOpts(client.WithVersion("1.45"))
		if err != nil {
			log.Printf("error creating docker client: %s", err.Error())
			panic(err)
		}

		testContainerName := fmt.Sprintf("test-%v", time.Now().UnixMilli())
		var containerID string
		
		JustBeforeEach(func() {
			resp, err := rawDockerClient.ContainerCreate(context.Background(), &container.Config{
				Image: "hello-world",	
			}, &container.HostConfig{}, &network.NetworkingConfig{}, nil, testContainerName)
			containerID = resp.ID
			Expect(err).NotTo(HaveOccurred())
		})

		JustAfterEach(func() {
			err := rawDockerClient.ContainerRemove(context.Background(), containerID, container.RemoveOptions{})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return a list of running containers", func() {
			containers, err := dockerClient.GetRunningContainers(context.Background())
			for _, container := range containers {
				log.Printf("container: %s\n", container)
			}
			Expect(err).NotTo(HaveOccurred())
			Expect(len(containers)).To(BeNumerically(">", 0))
		})
	})
})
