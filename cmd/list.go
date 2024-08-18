/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists existing tasks",
	Args:  cobra.ExactArgs(0),
	Long: `Lists tasks that have not been deleted using the 'delete' command.
  By default, 'list' will list ALL currently existing tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cmd.Help()
		} else {
			tasks := fetchTasksAsList()

			allVal, _ := cmd.Flags().GetBool("all")
			if allVal {
				createTable(tasks)
				return
			}

			completedVal, _ := cmd.Flags().GetBool("completed")
			if completedVal {
				createTable(filterCompletedTasks(tasks))
				return
			}

			incompletedVal, _ := cmd.Flags().GetBool("incompleted")
			if incompletedVal {
				createTable(filterIncompletedTasks(tasks))
				return
			}

			createTable(tasks)
		}
	},
}

var (
	all         bool
	completed   bool
	incompleted bool
	stats       bool
)

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

func filterCompletedTasks(tasks []Task) []Task {
	var completed []Task

	for _, task := range tasks {
		if task.Done {
			completed = append(completed, task)
		}
	}

	return completed
}

func filterIncompletedTasks(tasks []Task) []Task {
	var incompleted []Task

	for _, task := range tasks {
		if !task.Done {
			incompleted = append(incompleted, task)
		}
	}

	return incompleted
}
