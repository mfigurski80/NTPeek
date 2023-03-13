package render

func renderId(fields []interface{}, config renderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	for i, field := range fields {
		res[i] = field.(string)
	}
	return res, nil
}
