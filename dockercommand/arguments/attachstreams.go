package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

// AttachStreams attaches STDIN, STDOUT and STDERR to docker run call.
func AttachStreams(commandDef config.CommandDefinition, builder builder.Builder) error {
	_ = commandDef

	if isDetached, ok := commandDef.GetIsDetached(); isDetached && ok {
		return nil
	}

	if isInteractive, ok := commandDef.GetIsInteractive(); !isInteractive && ok {
		return nil
	}

	builder.
		AttachTo("STDIN").
		AttachTo("STDOUT").
		AttachTo("STDERR")

	return nil
}
