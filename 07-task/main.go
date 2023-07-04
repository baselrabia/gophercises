package main

import "task/cmd"

//
//$ taskModel
//taskModel is a CLI for managing your TODOs.
//
//Usage:
//taskModel [command]
//
//Available Commands:
//add         Add a new taskModel to your TODO list
//do          Mark a taskModel on your TODO list as complete
//list        List all of your incomplete tasks
//
//Use "taskModel [command] --help" for more information about a command.
//
//$ taskModel add review talk proposal
//Added "review talk proposal" to your taskModel list.
//
//$ taskModel add clean dishes
//Added "clean dishes" to your taskModel list.
//
//$ taskModel list
//You have the following tasks:
//1. review talk proposal
//2. some taskModel description
//
//$ taskModel do 1
//You have completed the "review talk proposal" taskModel.
//
//$ taskModel list
//You have the following tasks:
//1. some taskModel description

func main() {
	cmd.Execute()
}
