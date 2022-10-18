package main

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func formatRelativeDate(t time.Time) string {
	// convert "2022-07-01" to "next Monday"
	// t, _ := time.Parse("2006-01-02", date[:10])
	hoursDiff := time.Until(t).Hours()
	diff := math.Ceil((hoursDiff) / 24)
	// fmt.Printf("%v -> %v", hoursDiff/24, diff)
	if diff == 0 {
		return "Today"
	}
	if diff == 1 {
		return "Tomorrow"
	}
	if diff == -1 {
		return "Yesterday"
	}
	if diff < 10 && diff > -7 {
		if diff >= 8 {
			return "next " + t.Weekday().String()
		}
		if diff > 0 {
			return t.Weekday().String()
		}
		return "last " + t.Weekday().String()
	}
	if diff > 0 {
		return "in " + fmt.Sprint(diff) + " days"
	}
	return fmt.Sprint(-diff) + " days ago"
}

type TextHighlight struct {
	Bg   string
	Fore string
}

var colorMap = map[string]TextHighlight{
	"pink":    {"218", "0"},
	"red":     {"203", "0"},
	"orange":  {"208", "0"},
	"yellow":  {"219", "0"},
	"green":   {"120", "0"},
	"blue":    {"39", "0"},
	"purple":  {"141", "0"},
	"brown":   {"101", "15"},
	"gray":    {"248", "0"},
	"default": {"240", "15"},
}

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
		hi := colorMap[task.ClassColor]
		class := lipgloss.NewStyle().
			Background(lipgloss.Color(hi.Bg)).
			Foreground(lipgloss.Color(hi.Fore)).
			Render(task.Class)
		class = lipgloss.NewStyle().
			Width(maxClassLen).
			Align(lipgloss.Right).
			Render(class)

		name := lipgloss.NewStyle().
			Bold(true).
			Render(task.Name)

		relDate := fmt.Sprintf("(%s)", formatRelativeDate(task.Due))
		overdue := time.Until(task.Due).Hours()+24 < 0
		due := relDate
		if overdue {
			hi := colorMap["red"]
			due = lipgloss.NewStyle().
				Background(lipgloss.Color(hi.Bg)).
				Foreground(lipgloss.Color(hi.Fore)).
				Render(relDate)
		} else {
			due = lipgloss.NewStyle().
				Faint(!overdue).
				Render(relDate)
		}
		// fmt.Printf("%s %s %s\n", class, name, due)
		fmt.Printf("%s | %s  %s\n", class, name, due)
	}
}
