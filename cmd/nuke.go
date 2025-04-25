package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var nukeCmd = &cobra.Command{
	Use:   "nuke",
	Short: "Delete all your todos",
	Long:  "Destructive action that will remove all your todos from the list and reset your id counter.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This action will delete all your todos. Are you absolutely sure? (y/n)")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		answer := input.Text()
		if err := input.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if strings.ToLower(answer) != "y" {
			return
		}

		todoList.Nuke()

		write(cfg, todoList)

		fmt.Println("Todos deleted")
	},
}
