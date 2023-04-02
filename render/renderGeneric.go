package render

import (
	"fmt"

	"github.com/mfigurski80/NTPeek/priority"
	"golang.org/x/exp/maps"
)

type renderRowConfig struct {
	Name      string
	Modifiers []string
	Priority  *[]priority.Priority
}

type renderRowFunction func([]interface{}, renderRowConfig) ([]string, error)

var fRenderFuncs = map[string]renderRowFunction{
	"title":        renderTitle,
	"rich_text":    renderText,
	"select":       renderSelect,
	"multi_select": renderMultiSelect,
	"date":         renderDate,
	"checkbox":     renderCheckbox,
	"number":       renderNumber,
	"_id":          renderId,
	"_p":           renderPriority,
}

func renderNil(vals []interface{}, _ renderRowConfig) ([]string, error) {
	return make([]string, len(vals)), nil
}

// Figure out field type and return single common RenderRowFunction
func getGenericRenderFunc(field []interface{}, name string) (renderRowFunction, error) {
	fVals, ok := field[0].(map[string]interface{})
	if !ok {
		if f, ok := fRenderFuncs[name]; ok {
			return f, nil
		}
		return renderNil, nil
	}
	gMod := withGlobalModifiers
	f, ok := fRenderFuncs[fVals["type"].(string)]
	if !ok {
		return renderNil, fmt.Errorf(
			errType.UnsupportedType,
			name,
			fVals["type"].(string),
			maps.Keys(fRenderFuncs),
		)
	}
	return gMod(f), nil
}
