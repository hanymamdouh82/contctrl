package config

import (
	"errors"
	"fmt"
	"os"
	"path"
)

func load() ([]byte, error) {

	p, err := configPath()
	if err != nil {
		return nil, err
	}

	// try to load config file, if not found, create empty one and print location for user
	if _, err := os.Stat(p); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// file not found, will create one
			if e := create(p); e != nil {
				return nil, e
			}
		} else {
			// another error, may be permission error
			return nil, fmt.Errorf("failed to load or create config file")
		}
	}

	b, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func configPath() (string, error) {
	cp, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	p := path.Join(cp, CONFIG_FILE_NAME)
	return p, nil
}
