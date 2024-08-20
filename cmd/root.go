/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"os"
	"path"
	"sort"
	"strconv"

	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "donna",
	Short: "A simple CLI Todo list",
	Long:  `Keeps track of your tasks like Donna so that you can perform as efficiently as Harvey`,
}

var csvFileName string = "donna.csv"

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fetchTasksAsList() []Task {
	file, err := os.Open(getCsvPath())
	if err != nil {
		os.Create(getCsvPath())
	}

	defer file.Close()

	var tasks []Task

	rows, err := readCsvIgnoreHeaders(file)
	if err != nil {
		return tasks
	}

	for _, row := range rows {

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

func fetchTasksAsMap() map[int]Task {
	file, err := os.Open(getCsvPath())
	if err != nil {
		panic(err)
	}

	defer file.Close()

	m := make(map[int]Task)

	rows, err := readCsvIgnoreHeaders(file)
	if err != nil {
		return m
	}

	for _, row := range rows {

		id, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		completed, err := strconv.ParseBool(row[3])
		if err != nil {
			panic(err)
		}

		m[id] = Task{id, row[1], row[2], completed}
	}

	return m
}

func readCsvIgnoreHeaders(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return data, nil
	}

	return data[1:], nil
}

func writeTasksToCsv(tasks []Task) {
	csv_file, err := os.OpenFile(
		getCsvPath(),
		os.O_RDWR|os.O_TRUNC|os.O_CREATE,
		0755,
	)
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

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Id < tasks[j].Id
	})

	defer csv_file.Close()

	if err := gocsv.MarshalFile(&tasks, csv_file); err != nil {
		panic(err)
	}
}

func getCsvPath() string {
	homedir, _ := os.UserHomeDir()
	return path.Join(homedir, csvFileName)
}
