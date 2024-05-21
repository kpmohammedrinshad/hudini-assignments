package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	id          int
	description string
	isDone      bool
}

var tasks []Task

func addTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter description")
	descrp, _ := reader.ReadString('\n')
	tasks = append(tasks, Task{description: descrp, isDone: false})
}

func listTask() {
	fmt.Println("-----------------")
	if len(tasks) == 0 {
		fmt.Println("No task")
	}
	for i, task := range tasks {
		status := "pending"
		if task.isDone {
			status = "done"
		}
		fmt.Printf("%d %s [%s]\n", i+1, task.description, status)
	}
}
func markTaskAsDone() {
	listTask()
	if len(tasks) == 0 {
		return
	}
	fmt.Println("----------")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number to mark as done:")
	mark, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(mark))

	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("invalid task number")
		return
	}
	tasks[index-1].isDone = true
	fmt.Println("task marked as done")
}
func deleteTask() {
	listTask()
	if len(tasks) == 0 {
		return
	}
	fmt.Println("----------")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number to delete:")
	mark, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(mark))

	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("invalid task number")
		return
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
	fmt.Println("task is deleted")
}

// Main function to display the menu and handle user input
func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		// fmt.Println(choice, "\n")

		if err != nil {
			fmt.Println("Invalid choice, please try again.")
			fmt.Println(err.Error())
			continue
		}
		switch choice {
		case 1:
			addTask()
		case 2:
			listTask()
		case 3:
			markTaskAsDone()
		case 4:
			deleteTask()
		case 5:
			fmt.Println("existing..")
			return
		}
	}
}
