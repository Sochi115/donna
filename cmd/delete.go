/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Deletes task",
	Args:  cobra.MaximumNArgs(1),
	Long: `Deletes task with the given [ID].
  Deleted tasks can NOT be recovered`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if cmd.Flags().Lookup("all").Changed {
				deleteAllTasks()
				return
			}

			if cmd.Flags().Lookup("completed").Changed {
				saveIncompletedTasks()
				return
			}

			cmd.Help()
		} else {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("[ERROR] '%s' is not a valid number!\n\n", args[0])
				cmd.Help()
				return
			}
			deleteTaskById(id)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	var all bool
	var completed bool

	deleteCmd.Flags().BoolVarP(&all, "all", "a", false, "delete all tasks")
	deleteCmd.Flags().BoolVarP(&completed, "completed", "c", false, "delete all completed tasks")
}

func deleteTaskById(id int) {
	tasksMap := fetchTasksAsMap()

	_, ok := tasksMap[id]

	if !ok {
		fmt.Printf("Task with ID %d does not exist!\nRun `donna list` to view all tasks", id)
		return
	}

	delete(tasksMap, id)

	tasksList := make([]*Task, 0, len(tasksMap))

	for _, tasks := range tasksMap {
		tasksList = append(tasksList, tasks)
	}

	writeTasksToCsv(tasksList)

	fmt.Printf("Successfully deleted task %d", id)
}

func deleteAllTasks() {
	fmt.Print("Do you want to delete ALL tasks? [y/n]: ")
	var confirmation string
	fmt.Scanln(&confirmation)

	if strings.EqualFold(confirmation, "y") {
		writeTasksToCsv(make([]*Task, 0))
		fmt.Println("Successfully deleted all tasks")
		return
	}
	fmt.Println("Cancelled task")
}

func saveIncompletedTasks() {
	fmt.Print("Do you want to delete all completed tasks? [y/n]: ")
	var confirmation string
	fmt.Scanln(&confirmation)

	if strings.EqualFold(confirmation, "y") {
		incompletedTasks := getIncompletedTasks(fetchTasksAsList())
		writeTasksToCsv(incompletedTasks)
		fmt.Println("Successfully deleted all completed tasks")
		return
	}
	fmt.Println("Cancelled task")
}
