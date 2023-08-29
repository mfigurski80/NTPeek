package renderField

import "fmt"

/// Render number field

func RenderNumber(fields []interface{}, config RenderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	if len(config.Modifiers) > 0 {
		return res, fmt.Errorf(
			errType.UnsupportedMod, config.Name, "number", config.Modifiers[0],
			_SUPPORTED_GLOBAL_MODIFIERS,
		)
	}
	var gErr error
	for i, r := range fields {
		num, ok := r.(map[string]interface{})["number"].(float64)
		if !ok {
			res[i] = ""
			gErr = fmt.Errorf(errType.Internal, config.Name, r)
			continue
		}
		res[i] = fmt.Sprintf("%v", num)
	}
	return res, gErr
}
