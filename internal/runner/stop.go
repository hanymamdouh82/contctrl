package runner

import (
	"fmt"
	"os"
	"os/exec"
)

func StopProject(filePath string) error {

	if err := stop(filePath, []string{}); err != nil {
		return err
	}

	return nil
}

func StopStack(filePath string, services []string) error {

	if len(services) == 0 {
		return fmt.Errorf("cannot stop empty stack")
	}

	if err := stop(filePath, services); err != nil {
		return err
	}

	return nil
}

// stop docker compose for file and services
// if service array is empty, it will stop the all services in the compose file
// stack is governed by caller function
func stop(filePath string, services []string) error {

	// Get abs path for Docker binary
	binary, err := exec.LookPath("docker")
	if err != nil {
		return err
	}

	args := []string{"compose", "-f", filePath, "stop"}

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
