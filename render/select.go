package render

import "github.com/charmbracelet/lipgloss"

/// Render select field with color

func renderSelect(fields []interface{}, modifiers []string) []string {
	res := make([]string, len(fields))
	for i, r := range fields {
		body := r.(map[string]interface{})["select"].(map[string]interface{})
		value := body["value"].(string)
		color := colorMap[body["color"].(string)]
		res[i] = lipgloss.NewStyle().
			Background(lipgloss.Color(color.Bg)).
			Foreground(lipgloss.Color(color.Fg)).
			Render(value)
	}
	// TODO: support modifiers? global: right, center, left. local: no-color
	// TODO: support priority?
	return res
}
