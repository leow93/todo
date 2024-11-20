package cfg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

type Config struct {
	TodosFile string
}

func (c Config) Dir() string {
	return strings.TrimSuffix(c.TodosFile, "todos.json")
}

func configHome(home string) string {
	return fmt.Sprintf("%s/.config/todo", home)
}

func configLocation(configHome string) string {
	return fmt.Sprintf("%s/config.json", configHome)
}

func todosLocation(configHome string) string {
	return fmt.Sprintf("%s/todos.json", configHome)
}

func writeConfig(path string, cfg Config) error {
	bs, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(cfg.Dir(), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(path, bs, 0644)
}

func New() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("home not found: %s", err.Error())
	}
	cfgHome := configHome(home)
	cfgLocation := configLocation(cfgHome)
	todosLocation := todosLocation(cfgHome)

	content, err := os.ReadFile(cfgLocation)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			// write default config
			cfg := Config{
				TodosFile: todosLocation,
			}
			err := writeConfig(cfgLocation, cfg)
			return cfg, err

		}
		return Config{}, err
	}

	var cfg Config
	if err = json.Unmarshal(content, &cfg); err != nil {
		return Config{}, err
	}

	return Config{
		TodosFile: cfg.TodosFile,
	}, nil
}
