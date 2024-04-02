package main

import (
	"bufio"
	"fmt"
	"os"
)

var todolist = []string{}

// Create a new scanner to read from standard input
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for {
		var command string
		greeting()
		fmt.Println("Type the command or type 'help' to see the command list")
		fmt.Scan(&command)

		switch command {
		case "help":
			getListCommand()
		case "add":
			addToDoList()
		case "delete":
			deleteToDoList()
		case "update":
			updateToDoList()
		default:
			fmt.Println("Yor command is not available, check the available command list with type 'help'")
		}
	}

}

func greeting() {
	fmt.Println("\nWelcome to Golist!")
	fmt.Println("Simple Go App To Do List")
	fmt.Println("Your to do list:")
	for index, value := range todolist {
		fmt.Printf("%d. %s\n", index+1, value)
	}
}

func getListCommand() {
	commandList := "The commands are: \nhelp: get all commands \nadd: add new list  \ndelete: delete to do list \nupdate: update to do list"
	fmt.Println(commandList)

}

func addToDoList() {
	fmt.Println("Add new list:")
	scanner.Scan()
	newList := scanner.Text()

	todolist = append(todolist, newList)
}

func deleteToDoList() {
	var number uint
	fmt.Println("Type the number you want to delete:")
	fmt.Scan(&number)
	index := number - 1
	todolist = append(todolist[:index], todolist[index+1:]...)
	fmt.Println("todolist after deletion:", todolist)

}

func updateToDoList() {
	var number uint
	fmt.Println("Type the number you want to update:")
	fmt.Scan(&number)
	fmt.Println("Type the new value:")
	scanner.Scan()
	newValue := scanner.Text()
	index := number - 1
	todolist[index] = newValue
	fmt.Println("Updated todolist:", todolist)
}
