package render

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var loc, _ = time.LoadLocation("Local")

/// Render date field relative to local time

func renderDate(fields []interface{}, modifiers []string) []string {
	res := make([]string, len(fields))
	for i, field := range fields {
		date := field.(map[string]interface{})["date"].(map[string]interface{})["start"].(string)
		t, err := time.ParseInLocation("2006-01-02", date[:10], loc)
		if err != nil {
			panic(err)
		}
		relative := fmt.Sprintf("(%s)", stringifyDateAsRelative(t))
		if overdue(t) {
			hi := colorMap["red"]
			res[i] = lipgloss.NewStyle().
				Background(lipgloss.Color(hi.Bg)).
				Foreground(lipgloss.Color(hi.Fg)).
				Render(relative)
		} else {
			res[i] = lipgloss.NewStyle().Faint(true).Render(relative)
		}
	}
	// TODO: support modifiers?
	// TODO: support priority?
	return res
}

func overdue(t time.Time) bool {
	return time.Until(t).Hours() < -24
}

/// utility: build relative date string

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
