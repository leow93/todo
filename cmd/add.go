package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func parseDueDate(dd string) (*time.Time, error) {
	if dd == "" {
		return nil, nil
	}

	layouts := []string{time.DateOnly, time.DateTime}

	for _, layout := range layouts {
		d, err := time.Parse(layout, dd)
		if err == nil {
			return &d, nil
		}
	}

	return nil, fmt.Errorf("due date must be one of these formats: %s", strings.Join(layouts, ", "))
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a todo to the list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		due, err := parseDueDate(dueDate)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
		description := strings.Join(args, " ")
		todoList.Add(description, due)
		err = write(cfg, todoList)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var dueDate string

func init() {
	addCmd.Flags().StringVarP(&dueDate, "due", "d", "", "set a due date, e.g. 2000-01-01")
}
