package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
)

// Table of commands, use as cmdTable[function](params)
var cmdTable = map[Arg]func([]Arg){
	"add":    add,
	"ls":     ls,
	"check":  check,
	"clear":  clear,
	"rm":     rm,
	"rm-all": rmAll,
	"help":   help,
	"time":   time,
}

// Execute List of Command line args as a td command
func Execute(args []Arg) {
	action := args[0]
	params := args[1:]

	if !isValidAction(action) {
		fmt.Println("Invalid argument , see `td help` for details")
		// print help
		return
	}

	cmdTable[action](params)
}

// Returns if arg is a valid td action
func isValidAction(action Arg) bool {
	_, ok := cmdTable[action]

	return ok
}

func time(params []Arg) {

	// [ID] [TIME]
	tasks = loadTasks()

	id, err := strconv.Atoi(string(params[0]))

	if err != nil {
		fmt.Println("Could not convert string to id")
	}

	taskToTime := tasks[id]

	timer := NewCountDown(ParseTime(params[1:]))

	for !timer.IsFinished() {
		timer.WaitForSeconds(1)

		fmt.Printf("\r%s\033[K", timer.String())
	}
	fmt.Println()

	completed := timer.AskIfFinished(taskToTime)

	if completed {
		fmt.Println("Congrats :)")
		taskToTime.IsChecked = true
	} else {
		fmt.Println("Thats okay, perhaps try again with more time")
	}

	tasks[id] = taskToTime
	updateTasks(tasks)
	ls(params)

}

// Adds a td item to the list, with description of params
func add(params []Arg) {
	if len(params) == 0 {
		fmt.Println("Error: add must take arguments, see `td help` for details")
		return
	}

	_, ok := CreateTask(params)

	if !ok {
		fmt.Println("Error: Task could be created")
		return
	}

	ls(params)

}

// Lists all of the users td items
func ls(params []Arg) {

	tasks = loadTasks()

	sort.Slice(tasks[:], func(i, j int) bool {
		return tasks[i].Id < tasks[j].Id
	})

	//format print better
	for _, t := range tasks {
		fmt.Println(t.String())
	}
}

// Marks a td item as completed
func check(params []Arg) {
	if len(params) != 1 {
		fmt.Println("Error: too many arguments in 'check', see `td help` for details")
		return
	}

	id, err := strconv.Atoi(string(params[0]))

	if err != nil {
		fmt.Println("check must take an integer argument, see `td help` for details")
		return
	}

	tasks = loadTasks()

	foundTask := false

	for i := range tasks {

		if tasks[i].Id == id {
			foundTask = true
			tasks[i].IsChecked = !tasks[i].IsChecked
			break
		}
	}

	if !foundTask {
		fmt.Println("Error: could not find task")
		return
	}

	updateTasks(tasks)
	ls(params)
}

// Removes all checked td items from the lsit
func clear(params []Arg) {

	tasks = loadTasks()

	var uncheckedTasks []task

	for _, curTask := range tasks {
		if !curTask.IsChecked {
			uncheckedTasks = append(uncheckedTasks, curTask)
		}
	}

	updateTasks(uncheckedTasks)
	ls(params)
}

// Remove td item whose id is params[0]
func rm(params []Arg) {
	if len(params) != 1 {
		fmt.Println("Error: too many arguments in 'check', see `td help` for details")
		return
	}

	id, err := strconv.Atoi(string(params[0]))

	if err != nil {
		fmt.Println("rm must take an integer argument, see `td help` for details")
		return
	}

	tasks = loadTasks()

	var newTasks []task

	for _, curTask := range tasks {
		if curTask.Id == id {
			continue
		}

		newTasks = append(newTasks, curTask)
	}

	updateTasks(newTasks)
	ls(params)
}

// Deletes all td items
func rmAll(params []Arg) {
	if len(params) != 0 {
		fmt.Println("rm-all should not take params")
		return
	}

	var emptyTasks []task

	updateTasks(emptyTasks)
	ls(params)
}

func help(params []Arg) {

}
