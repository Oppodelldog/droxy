package dockerrun

import (
	"github.com/Oppodelldog/droxy/config"
	"github.com/sirupsen/logrus"
)

func prepareCommandLineArguments(commandDef *config.CommandDefinition, arguments []string) []string {

	for index, argVal := range arguments {
		replacement := getReplacement(commandDef, argVal)
		if replacement != nil {
			arguments[index] = *replacement
		}
	}

	return arguments
}

func getReplacement(commandDef *config.CommandDefinition, s string) *string {
	if replaceArgs, ok := commandDef.GetReplaceArgs(); ok {
		for _, replaceMapping := range replaceArgs {
			if len(replaceMapping) != 2 {
				logrus.Warnf("invalid argument replacement mapping '%v'. Replacement mapping must consist of 2 array entries.", replaceMapping)
				continue
			}
			if replaceMapping[0] == s {
				return &replaceMapping[1]
			}
		}
	}

	return nil
}
