package arguments

import (
	"fmt"
	"os"
	"regexp"

	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
)

// BuildPortsFromParams sets mappings of host ports to container ports.
func BuildPortsFromParams(commandDef *config.CommandDefinition, builder builder.Builder) error {
	if portsFromParams, ok := commandDef.GetPortsFromParams(); ok {
		return buildPortsFromParams(portsFromParams, builder)
	}

	return nil
}

func buildPortsFromParams(portRegEx []string, builder builder.Builder) error {
	for _, regEx := range portRegEx {
		if portValue, ok := extractPortFromArgs(regEx); ok {
			builder.AddPortMapping(fmt.Sprintf("%s:%s", portValue, portValue))
		}
	}

	return nil
}

func extractPortFromArgs(regEx string) (string, bool) {
	r := regexp.MustCompile(regEx)

	for _, arg := range os.Args {
		m := r.FindAllStringSubmatch(arg, -1)
		if m != nil {
			return m[0][1], true
		}
	}

	return "", false
}
