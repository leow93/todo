package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/leow93/todo/config"
	"github.com/leow93/todo/todos"
)

func Done(cfg *config.Config, t *todos.Todos, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(`"todo done" expects an integer id, e.g. "todo done 4"`)
		os.Exit(1)
	}
	done := t.MarkDone(id)
	if !done {
		fmt.Println("No task with that id")
		return
	}

	write(cfg, t)
}
