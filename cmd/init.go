package cmd

import (
	"github.com/leow93/todo/config"
	"github.com/leow93/todo/storage"
	"github.com/leow93/todo/todos"
)

func Init() (*config.Config, *todos.Todos, error) {
	cfg, err := config.Read()
	if err != nil {
		return nil, nil, err
	}

	ts, err := storage.Read(cfg)
	if err != nil {
		return nil, nil, err
	}

	return cfg, ts, nil
}
