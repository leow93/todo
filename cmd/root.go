package cmd

import (
	"fmt"
	"os"

	"github.com/leow93/todo/config"
	"github.com/leow93/todo/storage"
	"github.com/leow93/todo/todos"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo is a simple todo list manager",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var (
	cfg      *config.Config
	todoList *todos.Todos
)

func initConfig() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ts, err := storage.Read(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg = c
	todoList = ts
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(configCmd, listCmd, addCmd, doneCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
