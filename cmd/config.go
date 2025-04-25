package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage config",
	Long:  "Use these commands to show or edit your config",
}

var printConfigCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the config",
	Long:  "Prints the config as JSON",
	Run: func(cmd *cobra.Command, args []string) {
		bs, err := json.Marshal(cfg)
		if err != nil {
			fmt.Println(`invalid config file`)
			os.Exit(1)
		}

		fmt.Printf("%s\n", string(bs))
	},
}

var editConfigCommand = &cobra.Command{
	Use:   "edit",
	Short: "Edit your config",
	Long:  "Uses your $EDITOR to open the config",
	Run: func(_ *cobra.Command, args []string) {
		editor := os.Getenv("EDITOR")
		if editor == "" {
			fmt.Println(errors.New("no $EDITOR set"))
			os.Exit(1)

		}
		file := cfg.ConfigFile()
		cmd := exec.Command(editor, file)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	configCmd.AddCommand(printConfigCmd, editConfigCommand)
}
