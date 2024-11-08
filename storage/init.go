package storage

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"

	"github.com/leow93/todo/config"
	"github.com/leow93/todo/todos"
)

func noSuchFile(err error) bool {
	e, ok := err.(*fs.PathError)
	return ok && errors.Is(e.Unwrap(), fs.ErrNotExist)
}

func initialise(cfg config.Config) error {
	_, err := os.Stat(cfg.TodoFile)

	if err != nil && noSuchFile(err) {
		if err = os.MkdirAll(cfg.TodoDir, os.ModePerm); err != nil {
			return err
		}

		f, err := os.Create(cfg.TodoFile)
		if err != nil {
			return err
		}
		ts := todos.New()
		bs, err := json.Marshal(ts)
		if err != nil {
			return err
		}
		_, err = f.Write(bs)
		if err != nil {
			return err
		}

		err = f.Close()
		if err != nil {
			return err
		}
		return nil

	}

	return err
}

func Read(cfg config.Config) (*todos.Todos, error) {
	err := initialise(cfg)
	if err != nil {
		return nil, err
	}

	bs, err := os.ReadFile(cfg.TodoFile)
	if err != nil {
		return nil, err
	}

	var ts todos.Todos
	err = json.Unmarshal(bs, &ts)
	if err != nil {
		return nil, err
	}
	if ts.Entries == nil {
		var entries []todos.Entry
		ts.Entries = &entries
	}
	return &ts, err
}

func Write(cfg config.Config, data []byte) error {
	f, err := os.OpenFile(cfg.TodoFile, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	if err = f.Truncate(0); err != nil {
		return err
	}
	if _, err = f.Seek(0, 0); err != nil {
		return err
	}

	if _, err = f.Write(data); err != nil {
		return err
	}

	return f.Close()
}
