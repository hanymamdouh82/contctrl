package runner

import (
	"fmt"
	"os"
	"os/exec"
)

// Edit project compose file using $EDITOR defined in env
// Edit `contctrl.yml` file using $EDITOR defined in env
// config file exists as it is handled by config package in main.go
// No need to check if file is found here
//
// If no $EDITOR defined return with error
func Edit(filePath string) error {

	// Get $EDITOR from env
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return fmt.Errorf("EDITOR environment variable is not defined")
	}

	// Get abs path for Docker binary
	binary, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	args := []string{filePath}

	cmd := exec.Command(binary, args...)
	cmd.Stdout = os.Stdout // stream directly
	cmd.Stderr = os.Stderr // capture docker's actual output
	cmd.Stdin = os.Stdin   // needed if docker prompts for anything

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
