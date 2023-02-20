package render

import "github.com/mfigurski80/NTPeek/priority"

func renderCheckbox(fields []interface{}, modifiers []string, _ []priority.Priority) []string {
	res := make([]string, len(fields))
	for i, field := range fields {
		value := field.(map[string]interface{})["checkbox"].(bool)
		if value {
			res[i] = "✔️"
		} else {
			res[i] = "❌"
		}
	}
	return res
}
