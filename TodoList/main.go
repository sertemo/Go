package main

import (
	"fmt"
)

func main() {
	fmt.Println("##### Welcome to our TodoList app! #####")

	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch Nana's full course"
	var rewardDessert = "Reward myself with a pizza"
	var taskItems = []string{shortGolang, fullGolang, rewardDessert}

	printTasks(taskItems)
	fmt.Println()
	taskItems = addTask(taskItems, "Read Go book")
	printTasks(taskItems)
}

func printTasks(tasks []string) {
	fmt.Println("List of my ToDos:")
	for i, task := range tasks {
		fmt.Printf("%v: %v\n", i+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	updatedtaskItems := append(taskItems, newTask)
	return updatedtaskItems
}
