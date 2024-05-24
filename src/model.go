package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

var tasks []task

// Creates a task object in memory
func CreateTask(args []Arg) (*task, bool) {
	//write a new task to disk via saveTasks
	task := new(task)

	task.Id = getNextAvailableTaskNumber()

	task.Description = argsToDescription(args)
	task.IsChecked = false

	saveTask(task)

	return task, true
	//updateTasks()
}

// Sets the disk tasks to be the memory tasks passed in
func updateTasks(newTasks []task) {

	filename := "/var/tmp/tasks.json"

	err := os.Remove(filename)

	if err != nil {
		fmt.Println("Error: could not update tasks")
	}

	for _, task := range newTasks {
		saveTask(&task)
	}
}

// Gets the next available task number in the users td list
func getNextAvailableTaskNumber() int {

	tasks = loadTasks()

	sort.Slice(tasks[:], func(i, j int) bool {
		return tasks[i].Id < tasks[j].Id
	})

	nextNum := -1

	for i := 0; i < len(tasks); i++ {
		if i != tasks[i].Id {
			nextNum = i
			break
		}
	}

	if nextNum == -1 {
		nextNum = len(tasks)
	}

	return nextNum
}

// Loads the task from disk into memory
func loadTasks() []task {
	// loads task from disk
	filename := "/var/tmp/tasks.json"

	file, err := os.ReadFile(filename)

	if err != nil {
		// fmt.Println("Could not load", filename)
		createTaskFileIfNotExists()
	}

	json.Unmarshal(file, &tasks)

	return tasks
}

// Saves a specified task to the disk
func saveTask(t *task) {

	filename := "/var/tmp/tasks.json"

	createTaskFileIfNotExists()

	file, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Could not find", filename)
	}

	diskTasks := []task{}

	json.Unmarshal(file, &diskTasks)

	diskTasks = append(diskTasks, *t)

	dataBytes, e := json.Marshal(diskTasks)

	if e != nil {
		fmt.Println("Could not marshall data")
	}

	err = os.WriteFile(filename, dataBytes, 0644)

	if err != nil {
		fmt.Println("Could not write to file")
	}

}

// Creates a task.json file if it is not already present
func createTaskFileIfNotExists() {

	filename := "/var/tmp/tasks.json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
	}
}
