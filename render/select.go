package render

import (
	"github.com/charmbracelet/lipgloss"
)

/// Render select field with color

func renderSelect(fields []interface{}, modifiers []string) []string {
	res := make([]string, len(fields))
	// get maxlen
	maxLen := 0
	for _, r := range fields {
		body := r.(map[string]interface{})["select"].(map[string]interface{})
		value := body["name"].(string)
		if len(value) > maxLen {
			maxLen = len(value)
		}
	}
	alignStyle := lipgloss.NewStyle().Width(maxLen).Align(lipgloss.Right)
	// render
	for i, r := range fields {
		body := r.(map[string]interface{})["select"].(map[string]interface{})
		value := body["name"].(string)
		color := colorMap[body["color"].(string)]
		res[i] = lipgloss.NewStyle().
			Background(lipgloss.Color(color.Bg)).
			Foreground(lipgloss.Color(color.Fg)).
			Render(value)
		res[i] = alignStyle.Render(res[i])
	}
	// TODO: support modifiers? global: right, center, left. local: no-color
	// TODO: support priority?
	return res
}
