package docker

import (
	"context"
	"log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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
})
