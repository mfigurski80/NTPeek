package render

import "fmt"

func renderCheckbox(fields []interface{}, config renderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	if len(config.Modifiers) > 0 {
		return res, fmt.Errorf(
			errType.UnsupportedMod, config.Name, "checkbox", config.Modifiers[0],
			_SUPPORTED_GLOBAL_MODIFIERS,
		)
	}
	for i, field := range fields {
		value := field.(map[string]interface{})["checkbox"].(bool)
		if value {
			res[i] = "[x]"
		} else {
			res[i] = "[ ]"
		}
	}
	return res, nil
}
