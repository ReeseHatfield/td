package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	t "time"
)

type CountDown struct {
	seconds    int8
	minutes    int8
	hours      int8
	isFinished bool
}

func (countdown *CountDown) IsFinished() bool {
	return countdown.isFinished
}

func NewCountDown(hours, minutes, seconds int8) *CountDown {
	cd := new(CountDown)

	cd.isFinished = false

	cd.seconds = seconds
	cd.minutes = minutes
	cd.hours = hours

	return cd
}

func NewCountdownFromTime(t t.Time) *CountDown {
	cd := new(CountDown)

	cd.seconds = int8(t.Second())
	cd.minutes = int8(t.Minute())
	cd.hours = int8(t.Hour())

	return cd
}

func (cd CountDown) String() string {
	s := fmt.Sprintf("\r%01d:%02d:%02d", cd.hours, cd.minutes, cd.seconds)

	return s
}

func (cd *CountDown) WaitForSeconds(numSeconds int) {
	t.Sleep(t.Duration(numSeconds) * t.Second)

	cd.DecrementSeconds(numSeconds)
}

func (cd *CountDown) DecrementSeconds(numSeconds int) {
	cd.seconds -= int8(numSeconds)

	if cd.seconds < 0 {
		cd.seconds = 59
		cd.minutes -= 1
	}

	if cd.minutes < 0 {
		cd.minutes = 59
		cd.hours -= 1
	}

	if cd.hours == 00 && cd.minutes == 0 && cd.seconds == 0 {
		cd.Stop()
	}
}

func (cd *CountDown) Stop() {
	cd.isFinished = true
}

func AskIfFinished(t task) bool {

	fmt.Println("Timer expired!")
	fmt.Printf("Did you finish the task (%s) [y/n] ", t.Description)

	scanner := bufio.NewScanner(os.Stdin)

	var line string
	for scanner.Scan() {

		line = scanner.Text()

		if IsValidYesNoString(line) {
			break
		}
	}

	return rune(line[0]) == 'y'

}

// 15m -> (0, 15, 0)
// 1h 15m 4s -> (1, 15, 4)
func ParseTime(args []Arg) (hours, minutes, seconds int8) {
	for _, arg := range args {
		argStr := string(arg)

		unit := argStr[len(argStr)-1]
		value := argStr[:len(argStr)-1]

		v, _ := strconv.Atoi(value)

		switch unit {
		case 'h':
			hours += int8(v)
		case 'm':
			minutes += int8(v)
		case 's':
			seconds += int8(v)
		}
	}
	return
}
