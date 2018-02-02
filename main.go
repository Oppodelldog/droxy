package main

import (
	"docker-proxy-command/cmd"
	"os"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

func main() {

	var rootCmd = cmd.Root
	cmd.Root.AddCommand(cmd.Symlinks)

	if len(os.Args) >= 2 && os.Args[1] == "symlinks" {
		err := rootCmd.Execute()
		if err != nil {
			logrus.Info(err)
		}
	} else if len(os.Args) >= 1 && filepath.Base(os.Args[0]) == "docker-proxy" {
		rootCmd.Help()
	} else {
		cmd.ProxyDockerCommand()
	}
}
