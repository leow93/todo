package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	es := t.List()
	if len(es) == 0 {
		fmt.Println("Nothing todo :)")
		return
	}

	for _, e := range es {
		fmt.Printf("%s\nid: %d\nCreated at: %s\n=====\n", e.Text, e.ID, e.CreatedAt.Format(time.DateTime))
	}
}

func main() {
	switch os.Args[1] {
	case "add":
		txt := strings.Join(os.Args[2:], " ")
		t.Add(txt)
	case "list":
		list(t)
	case "done":
		inputId := os.Args[2]
		id, err := strconv.Atoi(inputId)
		if err != nil {
			panic(err)
		}
		err = t.MarkDone(id)
		if err != nil {
			panic(err)
		}
	case "nuke":
		err := t.Nuke()
		if err != nil {
			panic(err)
		}

	default:
		fmt.Println("expected `add` command")
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
