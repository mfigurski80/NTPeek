package main

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func stringifyDateAsRelative(t time.Time) string {
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

func formatRelativeDate(t time.Time) string {
	relDate := fmt.Sprintf("(%s)", stringifyDateAsRelative(t))
	overdue := !time.Now().Before(t)
	if overdue {
		hi := colorMap["red"]
		return lipgloss.NewStyle().
			Background(lipgloss.Color(hi.Bg)).
			Foreground(lipgloss.Color(hi.Fore)).
			Render(relDate)
	}
	return lipgloss.NewStyle().
		Faint(!overdue).
		Render(relDate)
}
