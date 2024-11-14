package config

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
	TodosFile  string
	configFile string
}

func (c Config) ConfigFile() string {
	return c.configFile
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

func writeConfig(path string, cfg Config) error {
	bs, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, bs, 0644)
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
			// write default config
			cfg := Config{
				TodosFile:  home + "/.config/todo/todos.json",
				configFile: cfgLocation,
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
		TodosFile:  cfg.TodosFile,
		configFile: cfgLocation,
	}, nil
}
