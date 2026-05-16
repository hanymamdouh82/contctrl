package catalog

import (
	"fmt"
	"os"
	"path"
	"strings"
)

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
