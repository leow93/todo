package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/leow93/todo/config"
	"github.com/leow93/todo/todos"
)

func Nuke(cfg *config.Config, t *todos.Todos) {
	fmt.Println("This action will delete all your todos. Are you absolutely sure? (y/n)")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	answer := input.Text()
	if err := input.Err(); err != nil {
		panic(err)
	}

	if strings.ToLower(answer) != "y" {
		return
	}

	t.Nuke()

	write(cfg, t)

	fmt.Println("Todos deleted")
}
