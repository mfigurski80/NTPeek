package render

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

/// Render select field with color

func renderSelect(fields []interface{}, config renderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	if len(config.Modifiers) > 0 {
		return res, fmt.Errorf(
			errType.UnsupportedMod, config.Name, "select", config.Modifiers[0],
			_SUPPORTED_GLOBAL_MODIFIERS,
		)
	}
	for i, r := range fields {
		body, ok := r.(map[string]interface{})["select"].(map[string]interface{})
		if !ok {
			res[i] = ""
			continue
		}
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
	return res, nil
}
