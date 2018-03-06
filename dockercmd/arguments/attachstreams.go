package arguments

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercmd/builder"
)

// AttachStreams attaches STDIN, STDOUT and STDERR to docker run call
func AttachStreams(commandDef *config.CommandDefinition, builder builder.Builder) error {
	_ = commandDef
	builder.
		AttachTo("STDIN").
		AttachTo("STDOUT").
		AttachTo("STDERR")

	return nil
}
