package renderField

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func RenderMultiSelect(fields []interface{}, config RenderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	if len(config.Modifiers) > 0 {
		return res, fmt.Errorf(
			errType.UnsupportedMod, config.Name, "multiselect", config.Modifiers[0],
			_SUPPORTED_GLOBAL_MODIFIERS,
		)
	}
	var gErr error
	for i, field := range fields {
		res[i] = ""
		ms, ok := field.(map[string]interface{})["multi_select"].([]interface{})
		if !ok {
			res[i] = ""
			gErr = fmt.Errorf(errType.Internal, config.Name, field)
			continue
		}
		for _, s := range ms {
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
	return res, gErr
}
