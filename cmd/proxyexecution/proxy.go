package proxyexecution

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/Oppodelldog/droxy/logger"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand"

	"github.com/Oppodelldog/droxy/logging"
)

const errorPreparingDockerCall = 900

type (
	// ConfigLoader loads configuration.
	ConfigLoader interface {
		Load() config.Configuration
	}
	// Builder builds a executable command object.
	CommandBuilder interface {
		BuildCommandFromConfig(commandDef config.CommandDefinition) (*exec.Cmd, error)
	}
	// CommandRunner runs a command.
	CommandRunner interface {
		RunCommand(cmd *exec.Cmd) error
	}
	// CommandResultHandler handles the result of an executed command.
	CommandResultHandler interface {
		HandleCommandResult(*exec.Cmd, error) int
	}
	// ExecutableNameParser parsed the name of the current executed file from cli arguments.
	ExecutableNameParser interface {
		ParseCommandNameFromCommandLine() string
	}
)

func ExecuteDroxyCommand(args []string) int {
	dockerRunCommandBuilder, err := dockercommand.NewBuilder()
	if err != nil {
		logger.Errorf("error preparing docker call: %v", err)

		return errorPreparingDockerCall
	}

	configLoader := config.NewLoader()
	commandResultHandler := newResultHandler()
	commandRunner := newCommandRunner()
	executableNameParser := newExecutableNameParser()

	return executeCommand(
		args,
		dockerRunCommandBuilder,
		configLoader,
		commandResultHandler,
		commandRunner,
		executableNameParser,
	)
}

// executeCommand executes a proxy command.
func executeCommand(
	args []string,
	commandBuilder CommandBuilder,
	configLoader ConfigLoader,
	commandResultHandler CommandResultHandler,
	commandRunner CommandRunner,
	executableNameParser ExecutableNameParser,
) int {
	cfg := configLoader.Load()
	cfg.Logging = true

	closeLogger := enableLogging(&cfg)
	defer closeLogger()

	logger.Infof("configuration load from: '%s'", cfg.GetConfigurationFilePath())
	logger.Info()

	logger.Infof("environment variables:")

	for _, envVar := range os.Environ() {
		logger.Info(envVar)
	}

	logger.Info("----------------------------------------------------------------------")

	logger.Infof("origin arguments:")

	for _, arg := range args {
		logger.Info(arg)
	}

	logger.Info("----------------------------------------------------------------------")

	commandName := executableNameParser.ParseCommandNameFromCommandLine()

	cmdDef, err := cfg.FindCommandByName(commandName)
	if err != nil {
		logger.Errorf("cannot find command definition for '%s', but got: %v", commandName, err)
	}

	cmd, err := commandBuilder.BuildCommandFromConfig(cmdDef)
	if err != nil {
		logger.Errorf("error preparing docker call for '%s': %v", commandName, err)

		return errorPreparingDockerCall
	}

	logger.Infof("calling docker to run '%s'", commandName)
	logger.Infof(strings.Join(cmd.Args, " "))
	err = commandRunner.RunCommand(cmd)

	exitCode := commandResultHandler.HandleCommandResult(cmd, err)

	return exitCode
}

func enableLogging(cfg *config.Configuration) (f func()) {
	f = func() {}

	if !cfg.Logging {
		logger.SetOutput(io.Discard)

		return
	}

	logfileWriter, err := logging.GetLogWriter(cfg)
	if err != nil {
		// no chance to log error output since running docker process has priority before logging
		logger.SetOutput(io.Discard)

		return
	}

	logger.SetOutput(logfileWriter)

	f = func() {
		err := logfileWriter.Close()
		if err != nil {
			logger.Error(err)
		}
	}

	return
}
