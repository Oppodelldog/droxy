package cmd

import (
	"docker-proxy-command/builder"
	"docker-proxy-command/config"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func ExecuteDockerCommand(cfg *config.Configuration) {

	commandName := filepath.Base(os.Args[0])
	cmd, err := builder.BuildCommandFromConfig(commandName, cfg)
	if err != nil {
		panic(err)
	}

	err = runCommand(cmd)

	if exitErr, ok := err.(*exec.ExitError); ok {

		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			os.Exit(status.ExitStatus())
		} else {
			log.Printf("Could not get exit code")
			os.Exit(990)
		}
	}

	if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
		os.Exit(status.ExitStatus())
	} else {
		log.Printf("Could not get exit code")
		os.Exit(991)
	}
}

func runCommand(cmd *exec.Cmd) error {

	return cmd.Run()
}
