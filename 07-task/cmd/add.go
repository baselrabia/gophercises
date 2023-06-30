package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"task/behelper"
	"task/data"
)

func init() {
	rootCmd.AddCommand(AddCmd)
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long:  "long desc for Add command, Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		AddExecute(cmd, args)
	},
}

func AddExecute(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		behelper.Exitf("Missing task details \n")
	}

	task := data.Task{Details: strings.Join(args, " ")}
	if err := data.CreateTask(&task); err != nil {
		behelper.Exitf("%v", err)
	}
	fmt.Printf("Added %q to your task list.\n", task.Details)
}
