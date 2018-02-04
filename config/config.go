package config

// Load loads the configuration file.
func Load() *Configuration {

	configFilePath, err := DiscoverConfigFile()
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
