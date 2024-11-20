package cmd

import (
	"github.com/leow93/todo/config"
	"github.com/leow93/todo/todos"
)

func Add(cfg *config.Config, t *todos.Todos, todo string) {
	t.Add(todo)
	write(cfg, t)
}
