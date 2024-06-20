package main

import "strings"

// Converts a slice of arguments to a string
func argsToDescription(args []Arg) string {
	str := ""

	for i := range args {
		str += string(args[i]) + " "
	}

	return str
}

func IsValidYesNoString(s string) bool {
	return strings.HasPrefix(s, "y") || strings.HasPrefix(s, "n")
}
