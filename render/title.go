package render

func renderTitle(fields []interface{}, modifiers []string) []string {
	res := make([]string, len(fields))
	for i, field := range fields {
		res[i] = parseNotionRichText(field.(map[string]interface{})["title"].([]interface{}))
	}
	// TODO: support modifiers?
	// TODO: support priority?
	return res
}

func parseNotionRichText(richText []interface{}) string {
	var text string
	for _, entry := range richText {
		text += entry.(map[string]interface{})["plain_text"].(string)
	}
	return text
}
