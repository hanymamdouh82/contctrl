package config

const (
	CONFIG_FILE_NAME = "contctrl.yml"
)

type Config struct {
	Sources []string `yaml:"sources"`
}
