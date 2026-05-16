package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ServiceCatalog struct {
	Name  string
	Image string
}

type DockerService struct {
	Image string `yaml:"image"`
}

type DockerComposeYaml struct {
	Services map[string]DockerService `yaml:"services" json:"services"`
}

// Reads yaml file and parse included services
// The primary goal is to define all services included into a yaml compose file
func Parse(filePath string) ([]ServiceCatalog, error) {

	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var ss DockerComposeYaml
	if err := yaml.Unmarshal(b, &ss); err != nil {
		return nil, err
	}

	ctlg := make([]ServiceCatalog, 0)
	for k, v := range ss.Services {
		sc := ServiceCatalog{
			Name:  k,
			Image: v.Image,
		}

		ctlg = append(ctlg, sc)
	}

	return ctlg, nil
}
