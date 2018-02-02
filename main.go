package main

import (
	"docker-proxy-command/cmd"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func main() {

	var rootCmd = cmd.Root
	cmd.Root.AddCommand(cmd.NewSymlinkCommand())
	cmd.Root.AddCommand(cmd.NewCloneCommand())

	if len(os.Args) >= 2 {
		switch os.Args[1] {

		case "symlinks":
			fallthrough
		case "clones":
			err := rootCmd.Execute()
			if err != nil {
				logrus.Info(err)
			}
		}
	} else if len(os.Args) >= 1 && filepath.Base(os.Args[0]) == "docker-proxy" {
		rootCmd.Help()
	} else {
		cmd.ProxyDockerCommand()
	}
}
