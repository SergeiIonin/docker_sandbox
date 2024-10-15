package docker

import (
	"context"
	"slices"
	"testing"

	"github.com/docker/docker/api/types/network"
)

func TestDockerClient(t *testing.T) {

	dc := NewDockerClient()
	netName := "test_network"
	net, err := dc.CreateNetwork(netName)

	t.Logf("Network created: %s", net)

	netsSum, _ := dc.apiClient.NetworkList(context.Background(), network.ListOptions{})

	nets := make([]string, 0, len(netsSum))
	for _, netSum := range netsSum {
		nets = append(nets, netSum.Name)
	}

	if !slices.Contains(nets, netName) {
		t.Errorf("Network %s wasn't created", netName)
	}

	if err != nil {
		t.Errorf("Error creating network: %v", err)
	}

	dc.apiClient.NetworkRemove(context.Background(), netName)
}
