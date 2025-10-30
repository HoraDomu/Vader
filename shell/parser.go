package shell

import "strings"

// ParseInput splits input into command and arguments
func ParseInput(input string) []string {
	// Simple split on whitespace; later can handle quotes
	return strings.Fields(input)
}
