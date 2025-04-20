package main

import (
	"fmt"
	"os"

	"github.com/q4ow/mygit/internal/repository"
)

func main() {
	gitCommand := os.Args[1]

	switch gitCommand {
	case "init":
		repository.InitRepository()
	case "add":
		repository.AddFiles(os.Args[2:])
	case "commit":
		message := "Default commit message"
		if len(os.Args) > 2 {
			message = os.Args[2]
		}
		repository.Commit(message)
	default:
		fmt.Printf("Unknown command: %s\n", gitCommand)
	}
}
