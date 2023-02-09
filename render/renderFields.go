package render

import (
	"fmt"
	"strings"

	"github.com/mfigurski80/NTPeek/types"
)

func getRenderedFields(tasks []types.NotionEntry, fields []string) [][]string {
	// parse each field: NAME[.MODIFIER]*
	fieldNames := make([]string, len(fields))
	fieldModifiers := make([][]string, len(fields))
	for i, field := range fields {
		fieldNames[i], fieldModifiers[i] = getFieldRenderDirective(field)
	}
	// get values list for each interesting field
	fieldVals := make([][]interface{}, len(fields))
	for i, name := range fieldNames {
		fieldVals[i] = make([]interface{}, len(tasks))
		for j, task := range tasks {
			fieldVals[i][j] = task[name]
		}
	}
	// render each field
	rendered := make([][]string, len(fields))
	for i, name := range fieldNames {
		renderFunc, ok := getFieldRenderFunc(fieldVals[i])
		if !ok {
			fmt.Println("ERROR formatting field", name)
			rendered[i] = make([]string, len(fieldVals[i]))
			continue
		}
		rendered[i] = renderFunc(fieldVals[i], fieldModifiers[i])
	}
	// return rendered fields
	return rendered
}

// Parse individual `"Class.R"` string into name and modifiers
func getFieldRenderDirective(s string) (string, []string) {
	parts := strings.Split(s, ".")
	return parts[0], parts[1:]
}

type RenderRowFunction func([]interface{}, []string) []string

func getFieldRenderFunc(field []interface{}) (RenderRowFunction, bool) {
	fVals, ok := field[0].(map[string]interface{})
	if !ok {
		return nil, false
	}
	switch fVals["type"].(string) {
	case "title":
		return renderTitle, true
	case "rich_text":
		return renderTitle, true
	case "select":
		return renderSelect, true
	case "date":
		return renderDate, true
	default:
		return func(d []interface{}, m []string) []string {
			return make([]string, len(d))
		}, false
	}
}
