package main

// Converts a slice of arguments to a string
func argsToDescription(args []Arg) string {
	str := ""

	for i := range args {
		str += string(args[i]) + " "
	}

	return str
}
