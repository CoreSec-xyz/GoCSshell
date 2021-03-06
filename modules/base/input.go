package base

import (
	"os"
	"os/exec"
	"strings"

	"github.com/CoreSec-xyz/GoCSshell/modules/commands"
	"github.com/CoreSec-xyz/GoCSshell/modules/external"
)

// Input verarbeitet build-in commands und exec System Programme
func Input(input string) error {

	// Split the input separate the command and the arguments.
	args := strings.Split(input, " ")

	for i, x := range args {
		args[i] = strings.Replace(x, " ", "", -1)
	}

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		err := commands.Cd(args)
		return err
	case "lazydocker":
		err := external.Lazydocker() // Only test if it work
		return err
	case "pwd":
		err := commands.Pwd()
		return err
	case "export":
		err := commands.Export(args)
		return err
	case "exit":
		os.Exit(0)
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
