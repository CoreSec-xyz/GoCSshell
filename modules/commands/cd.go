package commands

import (
	"os"
	"strings"
)

// Cd func fuer verzeichniswechsel
func Cd(args []string) error {

	// Finde User Home
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// 'cd' to home if path empty.
	if len(args) < 2 || args[1] == "" {
		//		return ErrNoPath
		return os.Chdir(home)
	}

	special := string([]rune(args[1]))[0]
	// Check for special path.
	switch string(special) {
	case "~": // beginnt im User Home
		path := strings.TrimLeft(args[1], "~")
		return os.Chdir(home + path)
	}
	// Change the directory and return the error.
	return os.Chdir(args[1])

}
