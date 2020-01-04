package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/CoreSec-xyz/GoCSshell/modules/base"
	"github.com/CoreSec-xyz/GoCSshell/modules/commands"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		commands.Pwd()
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
		base.Input(input)
	}
}
