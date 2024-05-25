package services

import (
	"testing"
)

func TestGetImages(t *testing.T) {
	ds := NewDockerImageService()
	images := ds.GetImages()
	if len(images) == 0 {
		t.Errorf("No images found")
	}
}