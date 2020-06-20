package dockercommand

import (
	"context"
	"fmt"

	"github.com/Oppodelldog/droxy/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func newDockerClientAdapter() (dockerClientAdapter, error) {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		err = fmt.Errorf("error building name argument, opening docker client failed: %w", err)
		return dockerClientAdapter{}, err
	}

	return dockerClientAdapter{dockerClient: dockerClient}, nil
}

type dockerClientAdapter struct {
	dockerClient *client.Client
}

func (a dockerClientAdapter) getAPIVersion() (string, error) {
	ctx := context.Background()

	v, err := a.dockerClient.ServerVersion(ctx)
	if err != nil {
		return "", err
	}

	return v.APIVersion, nil
}

func (a dockerClientAdapter) exists(containerName string) bool {
	ctx := context.Background()
	options := types.ContainerListOptions{
		All: true,
	}

	containers, err := a.dockerClient.ContainerList(ctx, options)
	if err != nil {
		logger.Errorf("error loading container list: %v", err)

		return false
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+containerName {
				return true
			}
		}
	}

	return false
}
