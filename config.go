// Package config used in config.yaml configuration file
package config

import (
	"github.com/spf13/viper"
)

var (
	config = &Config{}
)

// Config initial configuration for all projects
type Config struct {
	App
	Butent
	Database
}

// GetDB returns database configuration
func GetDB() *Database {
	return &config.Database
}

// Debug return is in debug mode, true or false
func (c *Config) Debug() bool {
	return c.App.Verbose
}

// LoadMain reading config data from configuration file and
// populiate it to Config
func LoadMain() (*Config, error) {
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// LoadCustom reading config data for custom sections by section name
func LoadCustom(name string, stru interface{}) error {
	if err := viper.Sub(name).Unmarshal(stru); err != nil {
		return err
	}
	return nil
}
