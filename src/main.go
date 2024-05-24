package main

import (
	"os"
)

func main() {
	args := ParseArgs(os.Args)

	Execute(args)

}

// Parses os.Args into type []Arg
func ParseArgs(args []string) (parsed []Arg) {

	if len(args) <= 1 {
		Help()
		os.Exit(1)
	}

	// remove first arg
	args = args[1:]

	if args[0] == "help" {
		Help()
		os.Exit(0)
	}

	for i := range args {
		parsed = append(parsed, Arg(args[i]))
	}

	return
}
