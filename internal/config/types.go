package config

import (
	"github.com/hanymamdouh82/contctrl/internal/runner"
)

const (
	CONFIG_FILE_NAME = "contctrl.yml"
)

type Config struct {
	Sources []string `yaml:"sources"`
}

func (c *Config) Edit() error {
	cfp, err := configPath()
	if err != nil {
		return err
	}

	if err := runner.Edit(cfp); err != nil {
		return err
	}

	return nil
}
