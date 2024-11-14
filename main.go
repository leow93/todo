package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/leow93/todo/config"
	"github.com/leow93/todo/storage"
	"github.com/leow93/todo/todos"
)

var (
	cfg config.Config
	t   *todos.Todos
)

func init() {
	c, err := config.Read()
	if err != nil {
		panic(err)
	}
	cfg = c

	ts, err := storage.Read(cfg)
	if err != nil {
		panic(err)
	}
	t = ts
}

func list(t *todos.Todos) {
	ts := t.List()
	if len(ts) == 0 {
		fmt.Println("Nothing to do :)")
		return
	}

	for _, e := range ts {
		fmt.Printf("%s\nid: %d\nCreated at: %s\n=====\n", e.Text, e.ID, e.CreatedAt.Format(time.DateTime))
	}
}

func editConfig() error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return errors.New("no $EDITOR")
	}
	file := cfg.ConfigFile()
	cmd := exec.Command(editor, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

const help = `
todo is a utility for managing a simple todo list. 

Usage:
  todo add <description>            Add an item to the list
  todo ls                           List items to be done 
  todo done <id>                    Mark an item as done 
  todo nuke                         Remove all items
  todo help                         Display this help
  todo config                       Display your current config
  todo config edit                  COMING SOON: Allows you to edit your config in your $EDITOR
`

func main() {
	switch os.Args[1] {
	case "add":
		txt := strings.Join(os.Args[2:], " ")
		t.Add(txt)
	case "ls":
		list(t)
	case "done":
		inputId := os.Args[2]
		id, err := strconv.Atoi(inputId)
		if err != nil {
			fmt.Println(`"todo done" expects an integer id, e.g. "todo done 4"`)
			os.Exit(1)
		}
		if !t.MarkDone(id) {
			fmt.Println("No task with that id")
		}
	case "nuke":
		t.Nuke()
	case "config":
		if len(os.Args) > 2 {
			edit := os.Args[2]
			if edit == "edit" {
				err := editConfig()
				if err != nil {
					panic(err)
				}
				return
			}
			fmt.Println("Unknown command")
			fmt.Print(help)
			os.Exit(1)
		}

		bs, err := json.Marshal(cfg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(bs))
	case "help":
		fmt.Print(help)
	default:
		fmt.Print(help)
		os.Exit(1)
	}

	data, err := json.Marshal(*t)
	if err != nil {
		panic(err)
	}
	err = storage.Write(cfg, data)
	if err != nil {
		panic(err)
	}
}
