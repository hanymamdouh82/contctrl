package runner

import (
	"fmt"
	"os"
	"os/exec"
)

// Restart stack services
// It will return if services is empty to avoid restarting all services
func RestartStack(filePath string, services []string) error {

	if len(services) == 0 {
		return fmt.Errorf("cannot pull empty stack")
	}

	if err := restart(filePath, services); err != nil {
		return err
	}

	return nil
}

// restart docker stack services as defined by metadata.yml file in catalog
// if service array is empty, it will restart the all services since catalog failback to `default`
// stack is governed by caller function
func restart(filePath string, services []string) error {

	// Get abs path for Docker binary
	binary, err := exec.LookPath("docker")
	if err != nil {
		return err
	}

	args := []string{"compose", "-f", filePath, "restart"}

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
