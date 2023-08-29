package renderField

import "fmt"

func RenderCheckbox(fields []interface{}, config RenderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	if len(config.Modifiers) > 0 {
		return res, fmt.Errorf(
			errType.UnsupportedMod, config.Name, "checkbox", config.Modifiers[0],
			_SUPPORTED_GLOBAL_MODIFIERS,
		)
	}
	var gErr error
	for i, field := range fields {
		value, ok := field.(map[string]interface{})["checkbox"].(bool)
		if !ok {
			res[i] = "   "
			gErr = fmt.Errorf(errType.Internal, config.Name, "checkbox", "bool", field)
			continue
		}
		if value {
			res[i] = "[x]"
		} else {
			res[i] = "[ ]"
		}
	}
	return res, gErr
}
