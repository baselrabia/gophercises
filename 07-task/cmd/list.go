package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"task/behelper"
	"task/data"
)

func init() {
	rootCmd.AddCommand(ListCmd)
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, _ []string) {
		tasks, err := data.ListTasks(false)
		if err != nil {
			behelper.Exitf("%v", err)
		}

		if len(tasks) == 0 {
			behelper.Exite("You don't have any incomplete tasks.")
		}

		fmt.Println("You have the following tasks:")
		for _, t := range tasks {
			fmt.Printf("%d. %s .  completed : %t \n", t.ID, t.Details, t.Completed)
		}
	},
}
