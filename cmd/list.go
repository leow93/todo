package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List things to do",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		ts := todoList.List()
		if len(ts) == 0 {
			fmt.Println("Nothing to do :)")
			return
		}

		for _, e := range ts {
			fmt.Printf("id: %d\n%s\n\n", e.ID, e.Text)
		}
	},
}
