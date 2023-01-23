package main

import (
	"fmt"
	"strings"
)

func filterTasksByIdRefs(tasks []Task, idRefs []string, idRefLen int) ([]Task, error) {
	// filter tasks to those that match taskIdRefs
	n := 0
	found := make([]bool, len(idRefs))
	for _, task := range tasks {
		for i, taskIdRef := range idRefs {
			if task.Id[:idRefLen] == taskIdRef {
				tasks[n] = task
				found[i] = true
				n++
				break
			}
		}
	}
	tasks = tasks[:n]

	// show error if any taskIdRefs were not found
	if len(idRefs) != n {
		missing := make([]string, len(idRefs)-n)
		for i, f := range found {
			if !f {
				missing[i] = idRefs[i]
			}
		}
		return tasks, fmt.Errorf("tasks not found: %s", strings.Join(missing, ", "))
	}

	return tasks, nil
}

func markNotionTasksDone(taskIdRefs []string) {
	// fetch tasks asynchronously
	tasksChannel := make(chan []Task)
	go func() {
		tasksChannel <- queryNotionTaskDB(NotionDatabaseId)
	}()

	// verify input real quick
	ln := len(taskIdRefs[0])
	for _, taskIdRef := range taskIdRefs {
		if len(taskIdRef) != ln {
			fmt.Println("Task IDs must be of the same length")
			return
		}
	}
	if ln < 2 {
		fmt.Println("Task IDs must be at least 2 characters long")
		return
	}

	// pull tasks back
	tasks := <-tasksChannel

	// filter tasks to those that match taskIdRefs
	tasks, err := filterTasksByIdRefs(tasks, taskIdRefs, ln)
	if err != nil {
		fmt.Println(err)
	}
	if len(tasks) == 0 {
		return
	}

	// mark tasks as done
	for _, task := range tasks {
		fmt.Printf("  %s", task.Id)
		err := mutateNotionMarkTaskDone(task.Id)
		if err != nil {
			fmt.Printf("\nFailed to update task: %s: %v\n", task.Id, err)
			return
		}
		fmt.Printf("\r✔️\n")
	}

}
