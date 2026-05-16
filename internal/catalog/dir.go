package catalog

import (
	"os"
	"path"

	"github.com/hanymamdouh82/contctrl/internal/metadata"
	"github.com/hanymamdouh82/contctrl/internal/parser"
)

// builds catalog for single source dir
func catalogDir(baseDir string) ([]Project, error) {

	ds, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}

	projects := make([]Project, 0)

	// Iterate over collected dirs to build initial Project objects
	for _, d := range ds {

		// Include dirs only
		if d.Type().IsDir() {
			i, err := d.Info()
			if err != nil {
				continue
			}

			// Exclude .git and similars
			if i.Name() != ".git" {
				p := Project{
					Name:    d.Name(),
					BaseDir: path.Join(baseDir, d.Name()),
				}

				// build list of yaml files only
				files, err := projectFiles(p)
				if err != nil {
					return nil, err
				}

				// read yaml contents to get services
				for i, f := range files {
					ss, err := parser.Parse(f.AbsPath)
					if err != nil {
						return nil, err
					}
					files[i].Services = ss
				}

				p.Files = files

				// Load project metadata
				var md metadata.Metadata
				md, err = metadata.Load(path.Join(p.BaseDir, METADATA))
				if err != nil {
					// we fall back here to build stack from all availables Services
					// We collect services from all files
					slist := []string{}
					for _, f := range p.Files {
						for _, s := range f.Services {
							slist = append(slist, s.Name)
						}
					}

					md = metadata.Metadata{
						Stacks: []metadata.Stack{
							{
								Name:     "default",
								Services: slist,
							},
						},
					}
				}

				p.Metadata = md
				projects = append(projects, p)
			}
		}
	}

	return projects, nil
}
