package config

import (
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Database struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Name string `yaml:"name"`
	} `yaml:"database"`
	GitPath []string `yaml:"git_paths"`
	CtLogs  []string `yaml:"ct_logs"`
}

func (c *Config) Load(env string) (*Config, error) {
	// If c is nil, allocate it so yaml.Unmarshal has a place to write
	if c == nil {
		c = &Config{}
	}

	yamlFile, err := os.ReadFile(env)
	if err != nil {
		return nil, err
	}

	// Now 'c' is a valid pointer to a Config struct
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
