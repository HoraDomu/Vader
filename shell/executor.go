package shell

import (
	"os"
	"os/exec"
)

// Execute runs external commands
func Execute(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
