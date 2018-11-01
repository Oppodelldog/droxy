package proxyexecution

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Oppodelldog/droxy/logging"
	"github.com/sirupsen/logrus"
)

func ExecuteDroxyCommand(args []string) int {
	dockerRunCommandBuilder := dockercommand.NewCommandBuilder()
	configLoader := config.NewLoader()
	commandResultHandler := NewCommandResultHandler()
	commandRunner := NewCommandRunner()
	executableNameParser := NewExecutableNameParser()

	return executeCommand(args, dockerRunCommandBuilder, configLoader, commandResultHandler, commandRunner, executableNameParser)
}

// executeCommand executes a proxy command
func executeCommand(args []string, commandBuilder CommandBuilder, configLoader ConfigLoader, commandResultHandler CommandResultHandler, commandRunner CommandRunner, executableNameParser ExecutableNameParser) int {

	cfg := configLoader.Load()
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
	for _, arg := range args {
		logrus.Info(arg)
	}
	logrus.Info("----------------------------------------------------------------------")

	commandName := executableNameParser.ParseCommandNameFromCommandLine()
	cmd, err := commandBuilder.BuildCommandFromConfig(commandName, cfg)
	if err != nil {
		logrus.Errorf("error preparing docker call for '%s': %v", commandName, err)
		return 900
	}
	logrus.Infof("calling docker ro tun '%s'", commandName)
	logrus.Infof(strings.Join(cmd.Args, " "))
	err = commandRunner.RunCommand(cmd)

	exitCode := commandResultHandler.HandleCommandResult(cmd, err)

	return exitCode
}
