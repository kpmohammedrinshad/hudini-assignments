package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// create a task struct
type Task struct {
	id          int
	description string
	isDone      bool
}

var todolist []Task
// create an add function with argument as slice of todolist of task
func addTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the description")
	descrp, _ := reader.ReadString('\n')
	todolist = append(todolist, Task{description: descrp, isDone: false})
}

func listTask(){
    if len(todolist)==0{
        fmt.Println("the task is empty")
    }
    for i,task:=range todolist{
        status:="pending"
        if task.isDone{
            status="Done"
        }
        fmt.Printf("%d %s [%s]\n",i+1,task.description, status)
    }

}
func markTaskAsDone(){
    listTask()
    if len(todolist)==0{
        return
    }
    reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the number")
	mark, _ := reader.ReadString('\n')
    index,err:=strconv.Atoi(strings.TrimSpace(mark))

    if err!=nil || index<1 || index>len(todolist){
        fmt.Println("invalid task number")
        return
    }
    todolist[index-1].isDone=true
    fmt.Println("task marked as done")
}
func deleteTask(){
    listTask()
    if len(todolist)==0{
        return
    }
    reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the number")
	mark, _ := reader.ReadString('\n')
    index,err:=strconv.Atoi(strings.TrimSpace(mark))

    if err!=nil || index<1 || index>len(todolist){
        fmt.Println("invalid task number")
        return
    }
    todolist=append(todolist[:index-1],todolist[index:]...)
    fmt.Println("the task is deleted")
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

		fmt.Println(choice, "\n")

		if err != nil {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		// Write code
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
            fmt.Println("existing")
            return 
	}
}
}