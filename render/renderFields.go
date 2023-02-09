package render

import (
	"fmt"
	"strings"

	"github.com/mfigurski80/NTPeek/types"
	"golang.org/x/exp/maps"
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
		if tasks[0][name] == nil {
			fmt.Printf("ERROR: field '%s' not found. Instead found: %v\n", name, maps.Keys(tasks[0]))
			continue
		}
		for j, task := range tasks {
			fieldVals[i][j] = task[name]
		}
	}
	// render each field
	rendered := make([][]string, len(fields))
	for i, name := range fieldNames {
		renderFunc, ok := getFieldRenderFunc(fieldVals[i])
		if !ok {
			fmt.Printf("ERROR: formatting field '%s'\n", name)
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
	case "multi_select":
		return renderMultiSelect, true
	case "date":
		return renderDate, true
	case "checkbox":
		return renderCheckbox, true
	default:
		fmt.Printf("ERROR: unsupported field type '%s'\n", fVals["type"].(string))
		return func(d []interface{}, m []string) []string {
			return make([]string, len(d))
		}, false
	}
}
