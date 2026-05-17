package catalog

import (
	"encoding/json"
	"fmt"
	"path"

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

func (p *Project) RestartStack() error {
	err := runner.RestartStack(p.SelectedFile.AbsPath, p.SelectedStack.Services)
	return err
}

func (p *Project) Edit() error {
	err := runner.Edit(p.SelectedFile.AbsPath)
	return err
}

func (p *Project) EditMeta() error {
	err := runner.Edit(path.Join(p.BaseDir, METADATA))
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
