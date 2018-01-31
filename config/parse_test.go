package config

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	yaml := `
Version="1"

[[command]]
	name = "php"
	image = "php:7.1.13"
	addGroups = true
	impersonate = true
	workDir = "/app"
	removeContainer=true
	volumes = [
		"${APP_DIR}:/app",
		"${SSH_AUTH_SOCK}:/run/ssh.sock",
		"/etc/passwd:/etc/passwd:ro",
		"/etc/group:/etc/group:ro",
		"/run/docker.sock:/run/docker.sock"
	]
	envvars = [
		"SSH_AUTH_SOCK:/run/ssh.sock",
		"DOCKER_HOST=unix:///run/docker.sock"
	]
`

	conf, err := parseFromBytes([]byte(yaml))
	if err != nil {
		t.Fatalf("Did not expect parseFromBytes to return an error but got: %v", err)
	}

	bytes, _ := json.Marshal(conf)
	fmt.Println(string(bytes))

	expected := &Configuration{
		Command: []CommandDefinition{
			{
				Name:  "php",
				Image: "php:7.1.13",
				Volumes: []string{
					"${APP_DIR}:/app",
					"${SSH_AUTH_SOCK}:/run/ssh.sock",
					"/etc/passwd:/etc/passwd:ro",
					"/etc/group:/etc/group:ro",
					"/run/docker.sock:/run/docker.sock",
				},
				EnvVars: []string{
					"SSH_AUTH_SOCK:/run/ssh.sock",
					"DOCKER_HOST=unix:///run/docker.sock",
				},
				AddGroups:       true,
				Impersonate:     true,
				WorkDir:         "/app",
				RemoveContainer: true,
			},
		},
		Version: "1",
	}

	assert.Equal(t, expected, conf)
}
