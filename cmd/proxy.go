package cmd

import (
	"docker-proxy-command/config"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"docker-proxy-command/logging"
	"io/ioutil"

	"docker-proxy-command/dockercmd"

	"github.com/sirupsen/logrus"
)

func ProxyDockerCommand() {

	cfg := config.Load()
	cfg.Logging = true
	if cfg.Logging {
		logfileWriter, err := logging.GetLogWriter(cfg)
		if err != nil {
			// no chance to log error output since running docker process has priority before logging
			logrus.SetOutput(ioutil.Discard)
		} else {
			logrus.SetOutput(logfileWriter)
			defer func() {
				err := logfileWriter.Close()
				if err != nil {
					logrus.Error(err)
				}
			}()
		}
	} else {
		logrus.SetOutput(ioutil.Discard)
	}

	logrus.Infof("configuration load from: '%s'", cfg.GetConfigurationFilePath())
	logrus.Info()

	logrus.Infof("environment variables:")
	for _, envVar := range os.Environ() {
		logrus.Info(envVar)
	}
	logrus.Info("----------------------------------------------------------------------")

	logrus.Infof("origin arguments:")
	for _, arg := range os.Args {
		logrus.Info(arg)
	}
	logrus.Info("----------------------------------------------------------------------")

	commandName := filepath.Base(os.Args[0])
	cmd, err := dockercmd.BuildCommandFromConfig(commandName, cfg)
	if err != nil {
		logrus.Errorf("error preparing docker call for '%s': %v", commandName, err)
		os.Exit(900)
	}
	logrus.Infof("calling docker ro tun '%s'", commandName)
	err = runCommand(cmd)

	if exitErr, ok := err.(*exec.ExitError); ok {

		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			logrus.Infof("docker finished with exit code '%v'", status.ExitStatus())
			os.Exit(status.ExitStatus())
		} else {
			logrus.Warning("Could not get exit code")
			os.Exit(990)
		}
	}

	if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
		logrus.Infof("docker finished with exit code '%v'", status.ExitStatus())
		os.Exit(status.ExitStatus())
	} else {
		logrus.Warning("Could not get exit code")
		os.Exit(991)
	}
}

func runCommand(cmd *exec.Cmd) error {

	return cmd.Run()
}
