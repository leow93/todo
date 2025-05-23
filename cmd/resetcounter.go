package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var resetCounterCmd = &cobra.Command{
	Use:   "reset-counter",
	Short: "Reduce counter to lowest possible value",
	Long:  "This command will mutate ids of existing todos such that your counter is now the length of your list.",
	Run: func(cmd *cobra.Command, args []string) {
		todoList.ResetCounter()
		write(cfg, todoList)
		fmt.Println("Counter reset")
	},
}
