package cmd

import (
	"fmt"
	"time"

	"github.com/leow93/todo/todos"
)

func List(t *todos.Todos) {
	ts := t.List()
	if len(ts) == 0 {
		fmt.Println("Nothing to do :)")
		return
	}

	for _, e := range ts {
		fmt.Printf("%s\nid: %d\nCreated at: %s\n=====\n", e.Text, e.ID, e.CreatedAt.Format(time.RFC1123Z))
	}
}
