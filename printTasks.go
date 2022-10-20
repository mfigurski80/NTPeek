package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	// mapset "github.com/deckarep/golang-set"
	"github.com/muesli/termenv"
)

func printTasks(tasks []Task) {
	lipgloss.SetColorProfile(termenv.TrueColor)
	maxClassLen := 0
	classLengths := make([]int, len(tasks))
	for i, task := range tasks {
		if len(task.Class) > maxClassLen {
			maxClassLen = len(task.Class)
		}
		classLengths[i] = len(task.Class)
	}
	for _, task := range tasks {
		// GET CLASS + FORMAT
		hi := colorMap[task.ClassColor]
		class := lipgloss.NewStyle().
			Background(lipgloss.Color(hi.Bg)).
			Foreground(lipgloss.Color(hi.Fore)).
			Render(task.Class)
		class = lipgloss.NewStyle().
			Width(maxClassLen).
			Align(lipgloss.Right).
			Render(class)

		// GET IMPORTANCE
		importanceVal := parseImportance(task)
		importance := formatImportance(
			importanceVal,
			[3]string{"│ ", "│ ", "│!"},
		)

		// GET TASK TEXT + FORMAT
		name := lipgloss.NewStyle().
			Bold(importanceVal != LO).
			Faint(importanceVal == LO).
			Render(task.Name)

		// GET DUE DATE + FORMAT
		due := formatRelativeDate(task.Due)

		// GET TASK ID
		id := lipgloss.NewStyle().
			Faint(true).
			Render(fmt.Sprintf("%.2s", task.Id))

		// PRINT
		// fmt.Printf("%s %s %s\n", class, name, due)
		fmt.Printf("%s %s%s  %s %s\n", class, importance, name, due, id)

	}
}
