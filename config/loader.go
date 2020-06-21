package config

import "github.com/Oppodelldog/filediscovery"

const configFileName = "droxy.toml"
const configEnvVarName = "DROXY_CONFIG"

//NewLoader returns a new config loader.
func NewLoader() Loader {
	return Loader{}
}

type (
	Loader struct{}
)

// Load loads the configuration file.
func (cl Loader) Load() Configuration {
	configFileDiscovery := createFileDiscovery()
	configFilePath, err := configFileDiscovery.Discover(configFileName)

	if err != nil {
		panic(err)
	}

	cfg, err := Parse(configFilePath)
	if err != nil {
		panic(err)
	}

	return cfg
}

func createFileDiscovery() filediscovery.FileDiscoverer {
	return filediscovery.New(
		[]filediscovery.FileLocationProvider{
			filediscovery.WorkingDirProvider(),
			filediscovery.ExecutableDirProvider(),
			filediscovery.EnvVarFilePathProvider(configEnvVarName),
		},
	)
}
