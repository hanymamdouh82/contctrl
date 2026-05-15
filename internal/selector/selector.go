package selector

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hanymamdouh82/contctrl/internal/navigator"
	"github.com/ktr0731/go-fuzzyfinder"
)

// Fuzzy find select project
func SelectProject(prjs []navigator.Project) navigator.Project {

	idx, err := fuzzyfinder.Find(prjs,
		func(i int) string {
			return prjs[i].Name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Project: %s ", prjs[i].Name)
		}))

	if err != nil {
		log.Fatal(err)
	}

	project := prjs[idx]

	return project
}

// Fuzzy find select file
func SelectFile(prj *navigator.Project) {
	fidx, err := fuzzyfinder.Find(prj.Files,
		func(i int) string {
			return prj.Files[i].Name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			var buff bytes.Buffer
			for _, s := range prj.Files[i].Services {
				buff.WriteString(s.Name + "\n")
			}
			return buff.String()
		}))

	if err != nil {
		log.Fatal(err)
	}

	prj.ActiveFile(fidx)

}

// Fuzzy find select stack
func SelectStack(prj *navigator.Project) {

	idx, err := fuzzyfinder.Find(prj.Metadata.Stacks,
		func(i int) string {
			return prj.Metadata.Stacks[i].Name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			var buff bytes.Buffer
			for _, s := range prj.Metadata.Stacks {
				buff.WriteString(s.Name + "\n")
			}
			return buff.String()
		}))

	if err != nil {
		log.Fatal(err)
	}

	prj.ActiveStack(idx)
}
