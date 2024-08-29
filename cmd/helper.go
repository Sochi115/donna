package cmd

func getCompletedTasks(tasks []*Task) []*Task {
	var completedTasks []*Task
	for _, task := range tasks {
		if task.Done {
			completedTasks = append(completedTasks, task)
		}
	}
	return completedTasks
}

func getIncompletedTasks(tasks []*Task) []*Task {
	var incompletedTasks []*Task
	for _, task := range tasks {
		if !task.Done {
			incompletedTasks = append(incompletedTasks, task)
		}
	}
	return incompletedTasks
}
