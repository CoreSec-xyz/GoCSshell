package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/CoreSec-xyz/GoCSshell/modules/base"
	"github.com/CoreSec-xyz/GoCSshell/modules/commands"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		if base.Config.ShowPwd == true {
			commands.Pwd()
		}
		fmt.Print(">")

		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err != nil {
			log.Fatal(err)
		}

		// Remove the newline character.
		input = strings.TrimSuffix(input, "\n")
		base.Input(input)
	}
}
