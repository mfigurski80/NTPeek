package render

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mfigurski80/NTPeek/priority"
)

var titleStyle = map[priority.Priority]lipgloss.Style{
	priority.HI: lipgloss.NewStyle().Bold(true),
	priority.LO: lipgloss.NewStyle().Faint(true),
}

func renderTitle(fields []interface{}, config renderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	for i, field := range fields {
		f, ok := field.(map[string]interface{})["title"].([]interface{})
		if !ok {
			res[i] = ""
			continue
		}
		res[i] = parseNotionRichText(f)
		switch config.Priority[i] {
		case priority.HI:
			res[i] = titleStyle[priority.HI].Render("!" + res[i])
		case priority.MED:
			res[i] = " " + res[i]
		case priority.LO:
			res[i] = titleStyle[priority.LO].Render(" " + res[i])
		}
	}
	// TODO: support modifiers?
	return res, nil
}

func parseNotionRichText(richText []interface{}) string {
	var text string
	for _, entry := range richText {
		text += entry.(map[string]interface{})["plain_text"].(string)
	}
	return text
}
