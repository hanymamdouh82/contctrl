package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hanymamdouh82/contctrl/internal/catalog"
	"github.com/hanymamdouh82/contctrl/internal/config"
	"github.com/hanymamdouh82/contctrl/internal/selector"
	"github.com/hanymamdouh82/contctrl/internal/ui"
)

func main() {
	// load config
	// If doesn't exist, config package will create it
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	// build projects catalog
	ps := catalog.Projects(&c)

	switch os.Args[1] {
	case "run":
		prj := prepProject(ps)
		run(prj)
	case "stop":
		prj := prepProject(ps)
		stop(prj)
	case "pull":
		prj := prepProject(ps)
		pull(prj)
	case "restart":
		prj := prepProject(ps)
		restart(prj)
	case "edit":
		prj := prepProject(ps)
		editCompose(prj)
	case "config":
		editConfig(&c)
	default:
		usage()
		os.Exit(1)
	}
}

// Opens fuzzy selection for projct and set default project file
// Invoke this function for commands that requires project selections
func prepProject(ps []catalog.Project) *catalog.Project {
	project, err := selector.SelectProject(ps)
	if err != nil {
		log.Fatal(err)
	}
	ui.Confirm("Project", project.Name)

	// We check files associated to project.
	// If more than one file, there is no default and we open select file
	// If only one file we assume it is the default file and use it as Activ File
	if len(project.Files) > 1 {
		if err := selector.SelectFile(&project); err != nil {
			log.Fatal(err)
		}
		ui.Confirm("File", project.GetActiveFile().Name)
	} else {
		project.ActiveFile(0)
	}

	return &project
}

// Command: run
func run(prj *catalog.Project) {
	if err := selector.SelectStack(prj); err != nil {
		log.Fatal(err)
	}
	ui.Confirm("Stack", prj.SelectedStack.Name)
	ui.Section("docker output")

	if err := prj.RunActiveStack(); err != nil {
		log.Fatal(err)
	}
}

// Command: stop
func stop(prj *catalog.Project) {
	ui.Section("docker output")
	if err := prj.StopProject(); err != nil {
		log.Fatal(err)
	}
}

// Command: pull
func pull(prj *catalog.Project) {
	if err := selector.SelectStack(prj); err != nil {
		log.Fatal(err)
	}
	ui.Confirm("Stack", prj.SelectedStack.Name)
	ui.Section("docker output")
	if err := prj.PullStack(); err != nil {
		log.Fatal(err)
	}
}

// Command: restart
func restart(prj *catalog.Project) {
	if err := selector.SelectStack(prj); err != nil {
		log.Fatal(err)
	}
	ui.Confirm("Stack", prj.SelectedStack.Name)
	ui.Section("docker output")
	if err := prj.RestartStack(); err != nil {
		log.Fatal(err)
	}
}

// Command: edit
func editCompose(prj *catalog.Project) {
	if err := prj.Edit(); err != nil {
		log.Fatal(err)
	}
}

// Command: config
func editConfig(c *config.Config) {
	if err := c.Edit(); err != nil {
		log.Fatal(err)
	}
}

// Prints CLI usage commands and verbs
func usage() {
	fmt.Println(`contctrl - Containerization Control Plane

Usage:
  contctrl run      Run stack services
  contctrl stop     Stop compose services
  contctrl restart  Restart a specific stack
  contctrl pull     Pull and restart a specific service
  contctrl edit     Edit a project compose file
  contctrl config   Edit contctrl config file`)
}
