package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PomodoroTimer(task task) {

	numSprints := 0

	for {
		numSprints += Sprint(task)

		breakMinutes := 5
		if numSprints%4 == 0 {
			breakMinutes = 20 // my personal preference
		}

		fmt.Print("Continue? [y/n]: ")

		scanner := bufio.NewScanner(os.Stdin)

		var line string
		for scanner.Scan() {

			line = scanner.Text()

			if IsValidYesNoString(line) {
				break
			}
		}

		if rune(line[0]) == 'n' {
			break
		}

		Break(task, breakMinutes)
	}

	completed := AskIfFinished(task)

	if completed {
		fmt.Println("Congrats :)")
		task.IsChecked = true
	} else {
		fmt.Println("Thats okay, perhaps try again with more time")
	}

}

func Sprint(task task) int {
	fmt.Printf("25 minute sprint for task: %s\n", task.Description)

	workTimer := NewCountDown(0, 25, 0)
	for !workTimer.IsFinished() {
		workTimer.WaitForSeconds(1)
		fmt.Printf("\r%s\033[K", workTimer.String())
	}
	fmt.Println()

	return 1
}

func Break(task task, minutes int) {

	fmt.Println("Sprint over: " + strconv.Itoa(minutes) + " minute break")
	breakTimer := NewCountDown(0, int8(minutes), 0)
	for !breakTimer.IsFinished() {
		breakTimer.WaitForSeconds(1)
		fmt.Printf("\r%s\033[K", breakTimer.String())
	}
	fmt.Println()

	fmt.Println("Break over")
}
