package config

type Config struct {
	Database struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"database"`
	GitPath []string `json:"git_paths"`
	CtLogs  []string `yaml:"ct_logs"`
}
