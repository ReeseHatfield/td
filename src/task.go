package main

import "fmt"

type task struct {
	Id          int
	Description string
	IsChecked   bool
}

func (t *task) String() string {

	checkChar := " "
	var color = "\033[34m" // blue

	var resetColor = "\033[0m"

	if t.IsChecked {
		checkChar = "X"
		color = "\033[32m" // green
	}

	return fmt.Sprintf("%s%d [%s] %s%s", color, t.Id, checkChar, t.Description, resetColor)
}
