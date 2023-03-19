package render

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func renderId(fields []interface{}, config renderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	// parse mod
	isShort := false
	for _, mod := range config.Modifiers {
		if mod == "short" {
			isShort = true
		} else {
			return nil, fmt.Errorf(errType.UnsupportedMod, config.Name, "id", mod, "[short]")
		}
	}
	// render result
	st := lipgloss.NewStyle().Faint(true)
	if isShort {
		for i, field := range fields {
			res[i] = st.Render(field.(string)[:4])
		}
	} else {
		for i, field := range fields {
			res[i] = st.Render(field.(string))
		}
	}
	return res, nil
}
