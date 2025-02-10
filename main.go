package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command>")
		return
	}

	switch os.Args[1] {
	case "add":
		// Handle add command
	case "list":
		// Handle list command
	case "delete":
		// Handle delete command
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}
