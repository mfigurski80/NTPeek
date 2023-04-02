package render

import "fmt"

/// Render number field

func renderNumber(fields []interface{}, config renderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	if len(config.Modifiers) > 0 {
		return res, fmt.Errorf(
			errType.UnsupportedMod, config.Name, "number", config.Modifiers[0],
			_SUPPORTED_GLOBAL_MODIFIERS,
		)
	}
	var gErr error
	for i, r := range fields {
		body, ok := r.(map[string]interface{})
		if !ok {
			res[i] = ""
			gErr = fmt.Errorf(errType.Internal, config.Name, r)
			continue
		}
		res[i] = fmt.Sprintf("%v", body["number"].(float64))
	}
	return res, gErr
}
