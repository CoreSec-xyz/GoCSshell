package commands

import (
	"errors"
	"os"
	"strings"
)

//Export Set Enviroment Variables
func Export(args []string) error {

	errNoEnv := errors.New("No valid env")

	if len(args) < 2 || args[1] == "" {
		return errNoEnv
	}
	env := strings.Split(args[1], "=")

	return os.Setenv(env[0], env[1])
}
