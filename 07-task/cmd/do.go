package cmd

import (
	"fmt"
	"strconv"
	"task/behelper"
	"task/data"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(DoCmd)
}

var DoCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			behelper.Exitf("Missing task ID\n")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			behelper.Exitf("%v\n", err)
		}
		task := data.Task{ID: id}
		if err := data.CompleteTask(&task); err != nil {
			behelper.Exitf("%v\n", err)
		}
		fmt.Printf("You have completed the %q task.\n", task.Details)
	},
}
