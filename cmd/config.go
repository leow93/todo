package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/leow93/todo/config"
)

func PrintConfig(c *config.Config) {
	bs, err := json.Marshal(c)
	if err != nil {
		fmt.Println(`invalid config file`)
		os.Exit(1)
	}

	fmt.Printf("%s\n", string(bs))
}

func EditConfig(c *config.Config) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return errors.New("no $EDITOR")
	}
	file := c.ConfigFile()
	cmd := exec.Command(editor, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
