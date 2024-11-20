package cmd

import (
	"github.com/leow93/todo/config"
	"github.com/leow93/todo/todos"
)

func Nuke(cfg *config.Config, t *todos.Todos) {
	t.Nuke()

	write(cfg, t)
}
