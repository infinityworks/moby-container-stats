package config

import cfg "github.com/infinityworks/go-common/config"

// Config - Defines application configuration
type Config struct {
	*cfg.BaseConfig
}

// Init - Initialises Config struct with safe defaults if not present
func Init() Config {

	ac := cfg.Init()

	appConfig := Config{
		&ac,
	}

	return appConfig
}
