package commands

import (
	"os"
)

// Cd func fuer verzeichniswechsel
func Cd(args []string) error {

	// 'cd' to home with empty path not yet supported.
	if len(args) < 2 || args[1] == "" {
		//		return ErrNoPath
		return os.Chdir("/")
	}

	// Check for special path.
	switch args[1] {
	case "~": // Funktioniert noch nicht mit ~/dir/path
		// Finde User Home
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		return os.Chdir(home)

	case "exit":
		os.Exit(0)
	}
	// Change the directory and return the error.
	return os.Chdir(args[1])

}
