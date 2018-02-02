package config

func Load() *Configuration {

	configFilePath, err := DiscoverConfigFile()
	if err != nil {
		panic(err)
	}

	cfg, err := Parse(configFilePath)
	if err != nil {
		panic(err)
	}

	return cfg
}
