package cmd

import (
	"fmt"

	"github.com/leow93/todo/todos"
)

func List(t *todos.Todos) {
	ts := t.List()
	if len(ts) == 0 {
		fmt.Println("Nothing to do :)")
		return
	}

	for _, e := range ts {
		fmt.Printf("id: %d\n%s\n\n", e.ID, e.Text)
	}
}
