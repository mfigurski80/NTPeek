package render

func renderCheckbox(fields []interface{}, _ renderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	for i, field := range fields {
		value := field.(map[string]interface{})["checkbox"].(bool)
		if value {
			res[i] = "✔️"
		} else {
			res[i] = "❌"
		}
	}
	return res, nil
}
