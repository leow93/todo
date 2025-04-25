package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/leow93/todo/remote"
	"github.com/spf13/cobra"
)

var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage remote storage",
	Long:  "Used for synchronising to a from a remote file",
}

var remoteSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Push local file to remote",
	Run: func(cmd *cobra.Command, args []string) {
		r := remote.New("localhost:21", "todo", "todo")

		data, err := json.Marshal(*todoList)
		if err != nil {
			fmt.Printf("fatal error: %s", err.Error())
			os.Exit(1)
		}

		err = r.Sync("remote_todos.json", data)
		if err != nil {
			fmt.Printf("fatal error: %s", err.Error())
			os.Exit(1)
		}
	},
}

var readRemoteCmd = &cobra.Command{
	Use:   "read",
	Short: "Read remote todos file",
	Run: func(cmd *cobra.Command, args []string) {
		r := remote.New("localhost:21", "todo", "todo")

		data, err := r.Read("remote_todos.json")
		if err != nil {
			fmt.Printf("fatal error: %s", err.Error())
			os.Exit(1)
		}

		fmt.Println(string(data))
	},
}

var pullRemoteCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull remote todos file and merge with local",
	Run: func(cmd *cobra.Command, args []string) {
		r := remote.New("localhost:21", "todo", "todo")

		data, err := r.Read("remote_todos.json")
		if err != nil {
			fmt.Printf("fatal error: %s", err.Error())
			os.Exit(1)
		}

		fmt.Println("GOT DATA. TODO: WRITE TO LOCAL FILE")
		fmt.Println(string(data))
	},
}

func init() {
	remoteCmd.AddCommand(remoteSyncCmd, readRemoteCmd, pullRemoteCmd)
}
