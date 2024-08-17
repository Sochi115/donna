/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

var csvPath string = "todo.csv"

var rootCmd = &cobra.Command{
	Use:   "donna",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fetchTasks() []Task {
	file, err := os.Open(csvPath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []Task

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range data[1:] {

		id, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		completed, err := strconv.ParseBool(row[3])
		if err != nil {
			panic(err)
		}

		tasks = append(tasks, Task{id, row[1], row[2], completed})
	}

	return tasks
}

func writeTasksToCsv(tasks []Task) {
	csv_file, err := os.OpenFile(csvPath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	err = csv_file.Truncate(0)
	if err != nil {
		panic(err)
	}

	_, err = csv_file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	defer csv_file.Close()

	if err := gocsv.MarshalFile(&tasks, csv_file); err != nil {
		panic(err)
	}
}
