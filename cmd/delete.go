/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Deletes task",
	Args:  cobra.ExactArgs(1),
	Long: `Deletes task with the given [ID].
  Deleted tasks can NOT be recovered`,
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
			deleteTask(id)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteTask(id int) {
	tasksMap := fetchTasksAsMap()

	_, ok := tasksMap[id]

	if !ok {
		fmt.Printf("Task with ID %d does not exist!\nRun `donna list` to view all tasks", id)
		return
	}

	delete(tasksMap, id)

	tasksList := make([]Task, 0, len(tasksMap))

	for _, tasks := range tasksMap {
		tasksList = append(tasksList, tasks)
	}

	writeTasksToCsv(tasksList)

	fmt.Printf("Successfully deleted task %d", id)
}

func generateCurrDateString() string {
	currTime := time.Now()

	return fmt.Sprintf("%d-%s-%d", currTime.Day(), currTime.Month().String(), currTime.Year())
}
