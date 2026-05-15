package navigator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/hanymamdouh82/contctrl/internal/metadata"
	"github.com/hanymamdouh82/contctrl/internal/parser"
	"github.com/hanymamdouh82/contctrl/internal/runner"
)

const (
	METADATA = "metadata.yml"
)

type CFile struct {
	Name     string
	AbsPath  string
	Services []parser.ServiceCatalog
}

type Project struct {
	Name          string
	BaseDir       string
	Files         []CFile
	SelectedFile  *CFile
	Metadata      metadata.Metadata
	SelectedStack *metadata.Stack
}

// Select file as active file to get executed, i is file index
func (p *Project) ActiveFile(i int) {
	p.SelectedFile = &p.Files[i]
}

func (p *Project) GetActiveFile() *CFile {
	return p.SelectedFile
}

// Select file as active file to get executed, i is file index
func (p *Project) ActiveStack(i int) {
	p.SelectedStack = &p.Metadata.Stacks[i]
}

func (p *Project) RunActiveStack() error {
	err := runner.RunStack(p.SelectedFile.AbsPath, p.SelectedStack.Services)
	return err
}

func (p *Project) StopProject() error {
	err := runner.StopProject(p.SelectedFile.AbsPath)
	return err
}

func (p *Project) PullStack() error {
	err := runner.PullStack(p.SelectedFile.AbsPath, p.SelectedStack.Services)
	return err
}

// For debugging and logging only
func (p *Project) Describe() {
	fmt.Printf("Project Name: %s\n", p.Name)

	fmt.Println("Project Files:")
	for _, f := range p.Files {
		fmt.Printf("==> %s\n", f.Name)

		for _, s := range f.Services {
			fmt.Printf("  --> %s\n", s.Name)
		}
	}

	fmt.Println("Project Metadata:")
	j, _ := json.MarshalIndent(p.Metadata, "", "  ")
	fmt.Printf("%s\n", j)
}

// Sacns base dir and identifies the projects
// The convection is to treat each dir as a project
// only dirs with yaml files are considered as projects
func Projects(baseDir string) []Project {

	ds, err := os.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
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
					log.Fatal(err)
				}

				// read yaml contents to get services
				for i, f := range files {
					ss, err := parser.Parse(f.AbsPath)
					if err != nil {
						log.Fatal(err)
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

	return projects
}

// scans each project and identifies yaml files
func projectFiles(ps Project) ([]CFile, error) {

	projectFiles := make([]CFile, 0)

	files, err := os.ReadDir(ps.BaseDir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		n := f.Name()

		// we depend on extension
		// We can read the file and confirm it is yaml structure so we can work with any file
		parts := strings.Split(n, ".")
		if len(parts) == 0 {
			return nil, fmt.Errorf("No files found with defined extension")
		}

		// if metadata file skip
		if parts[0] == "metadata" {
			continue
		}

		ext := parts[len(parts)-1]

		// check if ext is yml or yaml
		if ext == "yml" || ext == "yaml" {
			cf := CFile{
				Name:    n,
				AbsPath: path.Join(ps.BaseDir, n),
			}
			projectFiles = append(projectFiles, cf)
		}

	}

	return projectFiles, nil

}
