package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Run starts the shell loop
func Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		args := ParseInput(input) // parser.go

		if IsBuiltin(args) { // builtin.go
			if err := RunBuiltin(args); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		}

		if err := Execute(args); err != nil { // executor.go
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
