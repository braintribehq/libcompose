package network

import (
	"github.com/redradrat/libcompose/config"
	composeclient "github.com/redradrat/libcompose/docker/client"
	"github.com/redradrat/libcompose/project"
)

// DockerFactory implements project.NetworksFactory
type DockerFactory struct {
	ClientFactory composeclient.Factory
}

// Create implements project.NetworksFactory Create method.
// It creates a Networks (that implements project.Networks) from specified configurations.
func (f *DockerFactory) Create(projectName string, networkConfigs map[string]*config.NetworkConfig, serviceConfigs *config.ServiceConfigs, networkEnabled bool) (project.Networks, error) {
	cli := f.ClientFactory.Create(nil)
	return NetworksFromServices(cli, projectName, networkConfigs, serviceConfigs, networkEnabled)
}
