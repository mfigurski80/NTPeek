package render

import (
	"fmt"

	"github.com/mfigurski80/NTPeek/priority"
	f "github.com/mfigurski80/NTPeek/render/field"
	"golang.org/x/exp/maps"
)

type renderRowConfig struct {
	Name      string
	Modifiers []string
	Priority  *[]priority.Priority
}

var fRenderFuncs = map[string]f.RenderRowFunction{
	"title":        f.RenderTitle,
	"rich_text":    f.RenderText,
	"select":       f.RenderSelect,
	"multi_select": f.RenderMultiSelect,
	"date":         f.RenderDate,
	"checkbox":     f.RenderCheckbox,
	"number":       f.RenderNumber,
	"_id":          f.RenderId,
	"_p":           f.RenderPriority,
}

func renderNil(vals []interface{}, _ renderRowConfig) ([]string, error) {
	return make([]string, len(vals)), nil
}

// Figure out field type and return single common RenderRowFunction
func getGenericRenderFunc(field []interface{}, name string) (f.RenderRowFunction, error) {
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
