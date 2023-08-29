package renderField

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

/// Render select field with color

func RenderSelect(fields []interface{}, config RenderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	if len(config.Modifiers) > 0 {
		return res, fmt.Errorf(
			errType.UnsupportedMod, config.Name, "select", config.Modifiers[0],
			_SUPPORTED_GLOBAL_MODIFIERS,
		)
	}
	var gErr error
	for i, r := range fields {
		body, ok := r.(map[string]interface{})["select"].(map[string]interface{})
		if !ok {
			res[i] = ""
			gErr = fmt.Errorf(errType.Internal, config.Name, r)
			continue
		}
		value := body["name"].(string)
		color := colorMap[body["color"].(string)]
		res[i] = lipgloss.NewStyle().
			Background(lipgloss.Color(color.Bg)).
			Foreground(lipgloss.Color(color.Fg)).
			Render(value)
	}
	return res, gErr
}
