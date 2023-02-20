package render

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mfigurski80/NTPeek/priority"
)

/// Render select field with color

func renderSelect(fields []interface{}, modifiers []string, _ []priority.Priority) []string {
	res := make([]string, len(fields))
	for i, r := range fields {
		body := r.(map[string]interface{})["select"].(map[string]interface{})
		value := body["name"].(string)
		color := colorMap[body["color"].(string)]
		res[i] = lipgloss.NewStyle().
			Background(lipgloss.Color(color.Bg)).
			Foreground(lipgloss.Color(color.Fg)).
			Render(value)
		// res[i] = alignStyle.Render(res[i])
	}
	// TODO: support modifiers? global: right, center, left. local: no-color
	// TODO: support priority?
	return res
}
