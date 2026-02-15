package models

type Config struct {
	Database struct {
		Host string `yaml:"host"`
		Port string `json:"port"`
	} `yaml:"database"`
	CT struct {
	} `yaml:"ct_logs"`
}
