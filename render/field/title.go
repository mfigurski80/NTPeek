package renderField

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/mfigurski80/NTPeek/priority"
)

func RenderTitle(fields []interface{}, config RenderRowConfig) ([]string, error) {
	return renderGenericRichText(fields, config, title)
}

func RenderText(fields []interface{}, config RenderRowConfig) ([]string, error) {
	return renderGenericRichText(fields, config, text)
}

/// GENERIC RICH TEXT: for `rich_text`, `title` fields

type ntRichTextType string

const (
	text  ntRichTextType = "rich_text"
	title ntRichTextType = "title"
)

func renderGenericRichText(fields []interface{}, config RenderRowConfig, tp ntRichTextType) ([]string, error) {
	res := make([]string, len(fields))
	if len(config.Modifiers) > 0 {
		return res, fmt.Errorf(
			errType.UnsupportedMod, config.Name, tp, config.Modifiers[0],
			_SUPPORTED_GLOBAL_MODIFIERS,
		)
	}
	var gErr error
	for i, field := range fields {
		f, ok := field.(map[string]interface{})[string(tp)].([]interface{})
		if !ok {
			res[i] = ""
			gErr = fmt.Errorf(errType.Internal, config.Name, field)
			continue
		}
		res[i] = parseNotionRichText(f)
		switch (*config.Priority)[i] {
		case priority.HI:
			res[i] = titleStyle[priority.HI].Render(res[i])
		case priority.LO:
			res[i] = titleStyle[priority.LO].Render(res[i])
		}
	}
	return res, gErr
}

var titleStyle = map[priority.Priority]lipgloss.Style{
	priority.HI: lipgloss.NewStyle().Bold(true),
	priority.LO: lipgloss.NewStyle().Faint(true),
}

func parseNotionRichText(richText []interface{}) string {
	var text string
	for _, entry := range richText {
		text += entry.(map[string]interface{})["plain_text"].(string)
	}
	return text
}
