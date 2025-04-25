package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/leow93/todo/config"
	"github.com/leow93/todo/todos"
	"github.com/spf13/cobra"
)

func Add(cfg *config.Config, t *todos.Todos, todo string) {
	t.Add(todo)
	write(cfg, t)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a todo to the list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := strings.Join(args, " ")
		todoList.Add(description)
		err := write(cfg, todoList)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
