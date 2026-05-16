package catalog

import (
	"log"

	"github.com/hanymamdouh82/contctrl/internal/config"
)

// Sacns source dirs and identifies the projects
// The convection is to treat each dir as a project
// only dirs with yaml files are considered as projects
func Projects(config *config.Config) []Project {

	projects := make([]Project, 0)

	for _, s := range config.Sources {
		dps, err := catalogDir(s)
		if err != nil {
			log.Printf("cannot catalog dir: %s - skipped\n", s)
		}
		projects = append(projects, dps...)
	}

	return projects
}
