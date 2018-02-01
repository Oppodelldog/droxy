package main

import (
	"docker-proxy-command/cmd"
	"docker-proxy-command/config"
	"os"
	"fmt"
	"log"
)

func init() {
}

func main() {

	cfg := getConfig()

	if len(os.Args) == 2 && os.Args[1] == "symlinks" {
		fmt.Println("creating symlinks...")
		err := cmd.CreateSymlinks(cfg)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		cmd.ExecuteDockerCommand(cfg)
	}
}

func getConfig() *config.Configuration {

	configFilePath, err := cmd.DiscoverConfigFile()
	if err != nil {
		panic(err)
	}

	cfg, err := config.Parse(configFilePath)
	if err != nil {
		panic(err)
	}

	return cfg
}
