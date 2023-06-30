package cmd

import (
	"github.com/spf13/cobra"
	"task/behelper"
	"task/bolt"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a very fast static site generator",
}

func Execute() {
	// connect to db once
	db, err := bolt.Connect()
	if err != nil {
		behelper.Exitf("%v", err)
	}
	defer db.Close()

	if err := rootCmd.Execute(); err != nil {
		behelper.Exitf("%v", err)
	}

}
