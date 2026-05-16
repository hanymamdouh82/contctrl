package config

import (
	"gopkg.in/yaml.v3"
)

func Load() (Config, error) {

	b, err := load()
	if err != nil {
		return Config{}, err
	}

	c := Config{}
	yaml.Unmarshal(b, &c)

	return c, nil
}
