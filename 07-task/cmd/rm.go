package cmd

import (
	"fmt"
	"strconv"
	"task/behelper"
	"task/data/taskModel"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a taskModel on your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			behelper.Exitf("Missing taskModel ID\n")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			behelper.Exitf("%v\n", err)
		}
		task := taskModel.Task{ID: id}
		if err := taskModel.RemoveTask(&task); err != nil {
			behelper.Exitf("%v\n", err)
		}
		fmt.Printf("You have removed the %q taskModel.\n", task.Details)
	},
}
