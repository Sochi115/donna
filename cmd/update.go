/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var descFlag string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update [ID]",
	Short:   "Updates a task data",
	Aliases: []string{"u"},
	Long: `  Updates the data the task with the given [ID] 
  If no commands are passed, the description is updated by default.`,
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: false,

	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("[ERROR] '%s' is not a valid number!\n\n", args[0])
			cmd.Help()
			return
		}

		taskMap := fetchTasksAsMap()

		task, ok := taskMap[id]

		if !ok {
			fmt.Printf("Task with ID %d does not exist!\nRun `donna list` to view all tasks", id)
			return
		}

		if cmd.Flags().Lookup("desc").Changed {
			task.Desc = descFlag
		} else {
			return
		}

		taskMap[id] = task
		tasksList := make([]*Task, 0, len(taskMap))

		for _, tasks := range taskMap {
			tasksList = append(tasksList, tasks)
		}

		writeTasksToCsv(tasksList)

		fmt.Printf("Successfully updated task %d", id)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&descFlag, "desc", "d", "", "updates task description")
	updateCmd.MarkFlagRequired("desc")
}
