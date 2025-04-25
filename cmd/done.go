package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a todo as done",
	Long:  "Accepts the id of the todo as an argument",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		done := todoList.MarkDone(id)
		if !done {
			fmt.Println("No task with that id")
			return
		}

		write(cfg, todoList)
	},
}
