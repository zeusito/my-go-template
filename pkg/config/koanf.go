package config

import (
	"strings"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Configurations struct {
	Server ServerConfigurations `koanf:"server"`
}

type ServerConfigurations struct {
	Port string `koanf:"port"`
}

// LoadConfigurations Loads configurations depending upon the environment
func LoadConfigurations() (*Configurations, error) {
	k := koanf.New(".")
	err := k.Load(file.Provider("resources/config.toml"), toml.Parser())
	if err != nil {
		return nil, err
	}

	// Searches for env variables and will transform them into koanf format
	// e.g. SERVER_PORT variable will be server.port: value
	err = k.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil)
	if err != nil {
		return nil, err
	}

	var configuration Configurations

	err = k.Unmarshal("", &configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}
