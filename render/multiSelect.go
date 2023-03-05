package render

import (
	"github.com/charmbracelet/lipgloss"
)

func renderMultiSelect(fields []interface{}, _ renderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	for i, field := range fields {
		res[i] = ""
		for _, s := range field.(map[string]interface{})["multi_select"].([]interface{}) {
			body := s.(map[string]interface{})
			value := body["name"].(string)
			color := colorMap[body["color"].(string)]
			res[i] += lipgloss.NewStyle().
				Background(lipgloss.Color(color.Bg)).
				Foreground(lipgloss.Color(color.Fg)).
				Margin(0, 1, 0, 0).
				Render(value)
		}
		// remove last space
		if len(res[i]) > 0 {
			res[i] = res[i][:len(res[i])-1]
		}
	}
	return res, nil
}
