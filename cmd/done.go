/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [ID]",
	Short: "Sets task with ID 'completed' status to True",
	Args:  cobra.ExactArgs(1),
	Long: `Sets the 'Completed' status of the task with the given ID to True.
  Note: This command does NOT delete the task`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("[ERROR] '%s' is not a valid number!\n\n", args[0])
				cmd.Help()
				return
			}
			setTaskComplete(id)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func setTaskComplete(id int) {
	tasksMap := fetchTasksAsMap()

	completedTask, ok := tasksMap[id]

	if !ok {
		fmt.Printf("Task with ID %d does not exist!\nRun `donna list` to view all tasks", id)
		return
	}

	completedTask.Done = true

	tasksMap[id] = completedTask

	tasksList := make([]*Task, 0, len(tasksMap))

	for _, tasks := range tasksMap {
		tasksList = append(tasksList, tasks)
	}

	writeTasksToCsv(tasksList)

	fmt.Printf("Successfully set task %d as done", id)
}
