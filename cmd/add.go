/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add \"[Task]\"",
	Short: "Add a new task",
	Long: `Adds a new task list to the todo repository. 
  For example: donna add "Do something" adds a new task with the description "Do something"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		} else {
			addTask(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTask(desc string) {
	updatedTasks := appendNewTask(desc)

	writeTasksToCsv(updatedTasks)

	fmt.Printf("Successfully added new task '%s'", desc)
}

func appendNewTask(description string) []*Task {
	tasks := fetchTasksAsList()

	var newId int
	if len(tasks) == 0 {
		newId = 1
	} else {
		newId = tasks[len(tasks)-1].Id + 1
	}

	newTask := Task{
		newId,
		description,
		generateCurrDateString(),
		false,
	}

	tasks = append(tasks, &newTask)

	return tasks
}

func generateCurrDateString() string {
	currTime := time.Now()
	return fmt.Sprintf("%d-%s-%d", currTime.Day(), currTime.Month().String(), currTime.Year())
}
