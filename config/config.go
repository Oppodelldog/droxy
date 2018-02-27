package config

import "github.com/Oppodelldog/droxy/filediscovery"

const configFileName = "droxy.toml"
const configEnvVarName = "DROXY_CONFIG"

// Load loads the configuration file.
func Load() *Configuration {

	configFileDiscovery := createFileDiscovery()
	configFilePath, err := configFileDiscovery.Discover(configFileName)
	if err != nil {
		panic(err)
	}

	cfg, err := Parse(configFilePath)
	if err != nil {
		panic(err)
	}

	cfg.SetConfigurationFilePath(configFilePath)

	return cfg
}

func createFileDiscovery() filediscovery.FileDiscovery {
	return filediscovery.New(
		[]filediscovery.FileLocationProvider{
			filediscovery.WorkingDirProvider(),
			filediscovery.ExecutableDirProvider(),
			filediscovery.EnvVarFilePathProvider(configEnvVarName),
		},
	)
}
