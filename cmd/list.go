package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func printfln(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

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
			printfln("%s", e.Text)
			printfln("id: %d", e.ID)
			if e.Due != nil {
				d := *e.Due
				printfln("due: %s", d.Format(time.DateTime))
			}
			printfln("")
		}
	},
}
