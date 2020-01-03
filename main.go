package main

import (
	"bufio"
	"fmt"
	"gitlab.com/coresec.xyz/GoCSshell/modules/base"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = base.Input(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}