package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

type Config struct {
	TodoFile string
	TodoDir  string
}

func configHome(home string) string {
	return fmt.Sprintf("%s/.config/todo", home)
}

func configLocation(configHome string) string {
	return fmt.Sprintf("%s/config.json", configHome)
}

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("home not found: %s", err.Error())
	}
	cfgHome := configHome(home)
	cfgLocation := configLocation(cfgHome)

	content, err := os.ReadFile(cfgLocation)
	if err != nil {
		if e, ok := err.(*fs.PathError); ok && errors.Is(e.Unwrap(), fs.ErrNotExist) {
			return Config{
				TodoFile: home + "/.config/todo/todos.json",
				TodoDir:  cfgHome,
			}, nil
		}

		return Config{}, err
	}

	var cfg Config
	if err = json.Unmarshal(content, &cfg); err != nil {
		return Config{}, err
	}

	return Config{
		TodoFile: cfg.TodoFile,
		TodoDir:  cfgHome,
	}, nil
}
