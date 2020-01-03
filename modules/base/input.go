package base

import (
	"os"
	"os/exec"
	"strings"

	"gitlab.com/coresec.xyz/GoCSshell/modules/commands"
	"gitlab.com/coresec.xyz/GoCSshell/modules/external"
)

func Input(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input separate the command and the arguments.
	args := strings.Split(input, " ")

	for i, x := range args {
		args[i] = strings.Replace(x, " ", "", -1)
	}

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		return commands.Cd(args)
	case "lazydocker":
		return external.Lazydocker()
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
