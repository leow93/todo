package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/leow93/todo/cmd"
)

const help = `
todo is a utility for managing a simple todo list. 

Usage:
  todo add <description>            Add an item to the list
  todo ls                           List items to be done 
  todo done <id>                    Mark an item as done 
  todo nuke                         Remove all items
  todo help                         Display this help
  todo config                       Display your current config
  todo config edit                  Opens your config in your $EDITOR
`

func main() {
	c, t, err := cmd.Init()
	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		fmt.Print(help)
		return
	}

	switch os.Args[1] {
	case "add":
		txt := strings.Join(os.Args[2:], " ")
		cmd.Add(c, t, txt)
	case "ls":
		cmd.List(t)
	case "done":
		inputId := os.Args[2]
		cmd.Done(c, t, inputId)
	case "nuke":
		cmd.Nuke(c, t)
	case "config":
		if len(os.Args) > 2 {
			edit := os.Args[2]
			if edit == "edit" {
				cmd.EditConfig(c)
				return
			}
			fmt.Println("Unknown command")
			fmt.Print(help)
			os.Exit(1)
		}

		cmd.PrintConfig(c)
	case "help":
		fmt.Print(help)
	default:
		fmt.Println("Unknown command")
		fmt.Print(help)
		os.Exit(1)
	}
}
