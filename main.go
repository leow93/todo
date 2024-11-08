package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
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
	Entries *[]entry
}

func (db *database) addTodo(txt string) {
	id := db.Counter + 1
	db.Counter = id

	entries := *db.Entries

	entries = append(entries, entry{
		ID:        id,
		Text:      txt,
		CreatedAt: time.Now(),
	})
	db.Entries = &entries
}

func (db *database) list() []entry {
	return *db.Entries
}

func (db *database) markDone(id int) error {
	var entries []entry
	for _, e := range *db.Entries {
		if e.ID != id {
			entries = append(entries, e)
		}
	}

	db.Entries = &entries

	return nil
}

func (db *database) nuke() error {
	var entries []entry
	db.Entries = &entries
	return nil
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

	if err = f.Truncate(0); err != nil {
		return err
	}
	if _, err = f.Seek(0, 0); err != nil {
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
	emptyEntries := make([]entry, 0)
	bs, err := os.ReadFile("./todos.txt")
	if err != nil {
		return &database{
			Counter: 0,
			Entries: &emptyEntries,
		}, nil
	}

	db := database{
		Counter: 0,
		Entries: &emptyEntries,
	}
	// empty file
	if len(bs) == 0 {
		return &db, nil
	}
	err = json.Unmarshal(bs, &db)
	if db.Entries == nil {
		db.Entries = &emptyEntries
	}
	return &db, err
}

func list() error {
	db, err := readFile()
	if err != nil {
		return err
	}
	es := db.list()
	if len(es) == 0 {
		fmt.Println("Nothing todo :)")
		return nil
	}

	for _, e := range es {
		fmt.Printf("%s\nid: %d\nCreated at: %s\n=====\n", e.Text, e.ID, e.CreatedAt.Format(time.DateTime))
	}
	return nil
}

func nuke() error {
	db, err := readFile()
	if err != nil {
		return err
	}

	err = db.nuke()
	if err != nil {
		return err
	}
	return db.save()
}

func markDone(id int) error {
	db, err := readFile()
	if err != nil {
		return err
	}

	err = db.markDone(id)
	if err != nil {
		return err
	}

	return db.save()
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
	case "done":
		inputId := os.Args[2]
		id, err := strconv.Atoi(inputId)
		if err != nil {
			panic(err)
		}
		err = markDone(id)
		if err != nil {
			panic(err)
		}

	case "nuke":
		err := nuke()
		if err != nil {
			panic(err)
		}

	default:
		fmt.Println("expected `add` command")
		os.Exit(1)

	}
}
