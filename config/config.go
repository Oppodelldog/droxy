package config

import "github.com/Oppodelldog/filediscovery"

const configFileName = "droxy.toml"
const configEnvVarName = "DROXY_CONFIG"

//NewLoader returns a new config loader
func NewLoader() Loader {
	return &configLoader{}
}

type (
	// Loader is able to locate and load configuration from a config file
	Loader interface {
		Load() *Configuration
	}
	configLoader struct{}
)

// Load loads the configuration file.
func (cl *configLoader) Load() *Configuration {

	configFileDiscovery := cl.createFileDiscovery()
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

func (cl *configLoader) createFileDiscovery() filediscovery.FileDiscoverer {
	return filediscovery.New(
		[]filediscovery.FileLocationProvider{
			filediscovery.WorkingDirProvider(),
			filediscovery.ExecutableDirProvider(),
			filediscovery.EnvVarFilePathProvider(configEnvVarName),
		},
	)
}
