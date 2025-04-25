package cmd

import (
	"encoding/json"

	"github.com/leow93/todo/config"
	"github.com/leow93/todo/storage"
	"github.com/leow93/todo/todos"
)

func write(cfg *config.Config, t *todos.Todos) error {
	data, err := json.Marshal(*t)
	if err != nil {
		return err
	}
	err = storage.Write(cfg, data)
	if err != nil {
		return err
	}

	return nil
}
