/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	all         bool
	completed   bool
	incompleted bool
	stats       bool

	completedTasks   []Task
	incompletedTasks []Task
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Lists existing tasks",
	Args:    cobra.ExactArgs(0),
	Long: `Lists tasks that have not been deleted using the 'delete' command.
  By default, 'list' will list ALL currently existing tasks.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cmd.Help()
		} else {
			tasks := fetchTasksAsList()

			if len(tasks) == 0 {
				fmt.Println("Congratulations! There are no tasks to do")
				return
			}

			processTasks(tasks)

			if all {
				createTable(tasks)
				return
			}

			if completed {
				createTable(completedTasks)
				return
			}

			if incompleted {
				createTable(incompletedTasks)
				return
			}

			createTable(tasks)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "list all tasks")
	listCmd.Flags().BoolVarP(&completed, "completed", "c", false, "list completed tasks")
	listCmd.Flags().BoolVarP(&incompleted, "incompleted", "i", false, "list incompleted tasks")
	listCmd.Flags().BoolVarP(&stats, "stats", "s", false, "append statistics to footer of table")
}

func createTable(tasks []Task) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)

	t.AppendHeader(table.Row{"ID", "Description", "Created", "Completed"})

	var values []table.Row

	for _, task := range tasks {
		values = append(values, table.Row{task.Id, task.Desc, task.Created, task.Done})
	}

	t.AppendRows(values)
	t.AppendRow(table.Row{""})

	if stats {
		t.AppendSeparator()
		t.AppendFooter(table.Row{"", "", "TOTAL", len(tasks)})
	}

	t.Render()
}

func processTasks(tasks []Task) {
	for _, task := range tasks {
		if task.Done {
			completedTasks = append(completedTasks, task)
		} else {
			incompletedTasks = append(incompletedTasks, task)
		}
	}
}
