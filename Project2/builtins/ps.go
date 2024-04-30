package builtins

import (
	"fmt"
	"os"
	"os/exec"
)

func ProcessStatus() {
	cmd := exec.Command("ps", "-e")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running ps command: %v\n", err)
		os.Exit(1)
	}
}
