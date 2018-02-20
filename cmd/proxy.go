package cmd

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd"
	"github.com/Oppodelldog/droxy/helper"
	"github.com/Oppodelldog/droxy/logging"
	"github.com/sirupsen/logrus"
)

// ExecuteCommand executes a proxy command
func ExecuteCommand() {

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

	commandName := helper.ParseCommandNameFromCommandLine()
	cmd, err := dockercmd.BuildCommandFromConfig(commandName, cfg)
	if err != nil {
		logrus.Errorf("error preparing docker call for '%s': %v", commandName, err)
		os.Exit(900)
	}
	logrus.Infof("calling docker ro tun '%s'", commandName)
	logrus.Infof(strings.Join(cmd.Args, " "))
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

	cmd.Stdout = helper.NewLoggingWriter(os.Stdout, logrus.StandardLogger(), "StdOut")
	cmd.Stderr = helper.NewLoggingWriter(os.Stderr, logrus.StandardLogger(), "StdErr")
	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}
