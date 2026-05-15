package metadata

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Stack struct {
	Name     string
	Services []string
}

type Metadata struct {
	Stacks []Stack
}

func Load(filePath string) (Metadata, error) {

	b, err := os.ReadFile(filePath)
	if err != nil {
		return Metadata{}, err
	}

	var p map[string][]string
	if err := yaml.Unmarshal(b, &p); err != nil {
		log.Fatal(err)
	}

	if len(p) == 0 {
		return Metadata{}, fmt.Errorf("no metadata file found")
	}

	stacks := make([]Stack, 0)
	for k, v := range p {
		s := Stack{
			Name:     k,
			Services: v,
		}
		stacks = append(stacks, s)
	}

	return Metadata{
		Stacks: stacks,
	}, nil
}
