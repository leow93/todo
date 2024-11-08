package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	PathToFile string
}

var config Config

func readConfig() (Config, error) {
	return Config{
		PathToFile: "./todos.txt",
	}, nil
}

func touch() {
	f, err := os.OpenFile(config.PathToFile, os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func init() {
	cfg, err := readConfig()
	if err != nil {
		panic(err)
	}

	config = cfg

	if _, err = os.Stat(config.PathToFile); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			touch()
		}
	}
}

func addTodo(todo string) error {
	db, err := readFile()
	if err != nil {
		return err
	}
	db.addTodo(todo)
	return db.save()
}

type database struct {
	Counter int
	Entries []entry
}

func (db *database) addTodo(txt string) {
	id := db.Counter + 1
	db.Counter = id

	db.Entries = append(db.Entries, entry{
		ID:        id,
		Text:      txt,
		CreatedAt: time.Now(),
	})
}

func (db *database) list() []entry {
	return db.Entries
}

func (db *database) save() error {
	f, err := os.OpenFile(config.PathToFile, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	bs, err := json.Marshal(*db)
	if err != nil {
		return err
	}

	s, err := f.Stat()
	if err != nil {
		return err
	}

	err = f.Truncate(s.Size())
	if err != nil {
		return err
	}
	if _, err = f.Write(bs); err != nil {
		return err
	}

	return nil
}

type entry struct {
	ID        int
	Text      string
	CreatedAt time.Time
}

func readFile() (*database, error) {
	bs, err := os.ReadFile("./todos.txt")
	if err != nil {
		return &database{}, nil
	}

	var db database
	// empty file
	if len(bs) == 0 {
		return &db, nil
	}
	err = json.Unmarshal(bs, &db)
	return &db, err
}

func list() error {
	db, err := readFile()
	if err != nil {
		return err
	}

	for _, e := range db.list() {
		fmt.Printf("%s\nid: %d\nCreated at: %s\n-----------------------------\n", e.Text, e.ID, e.CreatedAt.Format(time.DateTime))
	}
	return nil
}

func main() {
	switch os.Args[1] {
	case "add":
		txt := strings.Join(os.Args[2:], " ")
		err := addTodo(txt)
		if err != nil {
			panic(err)
		}

	case "list":
		err := list()
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("expected `add` command")
		os.Exit(1)

	}
}
