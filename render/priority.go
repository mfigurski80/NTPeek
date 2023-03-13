package render

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mfigurski80/NTPeek/priority"
)

func renderPriority(_ []interface{}, config renderRowConfig) ([]string, error) {
	res := make([]string, len(config.Priority))
	style := lipgloss.NewStyle().Bold(true)
	for i, p := range config.Priority {
		switch p {
		case priority.HI:
			res[i] = style.Render("!")
		default:
			res[i] = " "
		}
	}
	return res, nil
}
