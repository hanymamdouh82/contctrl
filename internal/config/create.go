package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Create default config file in default config locatiob based on OS
// Currently I support Linux only - Contributors may add more
func create(path string) error {

	dc := Config{
		Sources: []string{
			"~/your_projects_dir",
		},
	}

	b, err := yaml.Marshal(dc)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, b, os.ModePerm); err != nil {
		return err
	}

	return nil
}
