package main

import (
	"fmt"
)

func main() {
	var command string
	fmt.Println("Welcome to Golist!")
	fmt.Println("Simple Go App To Do List")
	fmt.Println("Type 'help' to see the command list")
	fmt.Scan(&command)

	if command == "help" {
		getListCommand()
	}

	todolist := []string{}
	var newList string
	fmt.Println("Add new list:")
	fmt.Scan(&newList)

	todolist = append(todolist, newList)
	fmt.Printf("All to do list: %v", todolist)

}

func getListCommand() {
	fmt.Println("Hi")
}
