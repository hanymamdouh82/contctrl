package runner

import (
	"fmt"
	"os"
	"os/exec"
)

// Pull stack services
// It will return if services is empty to avoid running all services
func PullStack(filePath string, services []string) error {

	if len(services) == 0 {
		return fmt.Errorf("cannot pull empty stack")
	}

	if err := pull(filePath, services); err != nil {
		return err
	}

	return nil
}

// pull image for docker compose file services
// if service array is empty, it will pull the all services in the compose file
// stack is governed by caller function
func pull(filePath string, services []string) error {

	// Get abs path for Docker binary
	binary, err := exec.LookPath("docker")
	if err != nil {
		return err
	}

	args := []string{"compose", "-f", filePath, "pull"}

	// Add stack services to arguments if any
	args = append(args, services...)

	cmd := exec.Command(binary, args...)
	cmd.Stdout = os.Stdout // stream directly
	cmd.Stderr = os.Stderr // capture docker's actual output
	cmd.Stdin = os.Stdin   // needed if docker prompts for anything

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
