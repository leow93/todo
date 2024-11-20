package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/leow93/todo/config"
	"github.com/leow93/todo/storage"
	"github.com/leow93/todo/todos"
)

func write(cfg *config.Config, t *todos.Todos) {
	data, err := json.Marshal(*t)
	if err != nil {
		fmt.Printf("fatal error: %s", err.Error())
		os.Exit(1)
	}
	err = storage.Write(cfg, data)
	if err != nil {
		fmt.Printf("fatal error: %s", err.Error())
		os.Exit(1)
	}
}
