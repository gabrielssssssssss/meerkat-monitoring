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
	if c == nil {
		c = &Config{}
	}

	yamlFile, err := os.ReadFile(env)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
