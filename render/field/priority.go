package renderField

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/mfigurski80/NTPeek/priority"
)

func RenderPriority(_ []interface{}, config RenderRowConfig) ([]string, error) {
	res := make([]string, len(*config.Priority))
	if len(config.Modifiers) > 0 {
		return res, fmt.Errorf(
			errType.UnsupportedMod, config.Name, "_p", config.Modifiers[0],
			_SUPPORTED_GLOBAL_MODIFIERS,
		)
	}
	style := lipgloss.NewStyle().Bold(true)
	for i, p := range *config.Priority {
		switch p {
		case priority.HI:
			res[i] = style.Render("!")
		default:
			res[i] = " "
		}
	}
	return res, nil
}
