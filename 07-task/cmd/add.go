package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"task/behelper"
	"task/data/taskModel"
)

func init() {
	rootCmd.AddCommand(AddCmd)
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new taskModel to your TODO list",
	Long:  "long desc for Add command, Add a new taskModel to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		AddExecute(cmd, args)
	},
}

func AddExecute(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		behelper.Exitf("Missing taskModel details \n")
	}

	task := taskModel.Task{Details: strings.Join(args, " ")}
	if err := taskModel.CreateTask(&task); err != nil {
		behelper.Exitf("%v", err)
	}
	fmt.Printf("Added %q to your taskModel list.\n", task.Details)
}
