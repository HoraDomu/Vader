package shell

import (
	"errors"
	"fmt"
	"os"
)

var ErrNoPath = errors.New("path required")

// Map of built-in commands
var builtins = map[string]func(args []string) error{
	"cd": func(args []string) error {
		if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])
	},
	"exit": func(args []string) error {
		os.Exit(0)
		return nil
	},
	"pwd": func(args []string) error {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(dir)
		return nil
	},
}

// IsBuiltin checks if a command is built-in
func IsBuiltin(args []string) bool {
	_, ok := builtins[args[0]]
	return ok
}

// RunBuiltin runs a built-in command
func RunBuiltin(args []string) error {
	return builtins[args[0]](args)
}
