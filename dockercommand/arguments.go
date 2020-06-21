package dockercommand

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/Oppodelldog/droxy/dockercommand/builder"
	"github.com/Oppodelldog/droxy/logger"
)

func prependAdditionalArguments(commandDef config.CommandDefinition, arguments []string) []string {
	if additionalArguments, ok := commandDef.GetAdditionalArgs(); ok {
		arguments = append(additionalArguments, arguments...)
	}

	return arguments
}

func buildArgumentsFromFunctions(
	commandDef config.CommandDefinition,
	builder builder.Builder,
	builders []argumentBuilderFunc,
) error {
	for _, argumentBuilderFunc := range builders {
		err := argumentBuilderFunc(commandDef, builder)
		if err != nil {
			return err
		}
	}

	return nil
}

func prepareCommandLineArguments(commandDef config.CommandDefinition, arguments []string) []string {
	for index, argVal := range arguments {
		replacement, ok := getReplacement(commandDef, argVal)
		if ok {
			arguments[index] = replacement
		}
	}

	return arguments
}

func getReplacement(commandDef config.CommandDefinition, s string) (string, bool) {
	if replaceArgs, ok := commandDef.GetReplaceArgs(); ok {
		for _, replaceMapping := range replaceArgs {
			const mustHaveEntries = 2
			if len(replaceMapping) != mustHaveEntries {
				logger.Warnf(
					"invalid argument replacement mapping '%v'. Replacement mapping must consist of 2 array entries.",
					replaceMapping,
				)

				continue
			}

			if replaceMapping[0] == s {
				return replaceMapping[1], true
			}
		}
	}

	return "", false
}
