package dockercommand

import "github.com/Oppodelldog/droxy/config"

func prependAdditionalArguments(commandDef *config.CommandDefinition, arguments []string) []string {
	if additionalArguments, ok := commandDef.GetAdditionalArgs(); ok {
		arguments = append(additionalArguments, arguments...)
	}

	return arguments
}
