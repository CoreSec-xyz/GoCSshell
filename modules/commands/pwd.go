package commands

import (
	"fmt"
	"os"
)

// Pwd gibt das Aktuelle verzeichnis zurueck
func Pwd() error {

	currentpwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println(currentpwd)

	return nil // Das Error handling fehlt noch... mache ich wenn ich es verstanden habe :-)
}
