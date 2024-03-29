package render

import (
	"github.com/acarl005/stripansi"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/maps"
)

/// Global modifiers: can be applied to ANY field type because...
/// it's applied to a final render, not a raw field value

// Existing global style + info on what we're styling => new global style
type modifierFunc func(lipgloss.Style, []string) lipgloss.Style

var _GLOBAL_RENDER_MODIFIERS map[string]modifierFunc = map[string]modifierFunc{
	"bold": func(s lipgloss.Style, _ []string) lipgloss.Style {
		return s.Bold(true)
	},
	"right": func(s lipgloss.Style, data []string) lipgloss.Style {
		maxLen := getMaxLen(data)
		return s.Width(maxLen).Align(lipgloss.Right)
	},
	"left": func(s lipgloss.Style, data []string) lipgloss.Style {
		maxLen := getMaxLen(data)
		return s.Width(maxLen).Align(lipgloss.Left)
	},
	"center": func(s lipgloss.Style, data []string) lipgloss.Style {
		maxLen := getMaxLen(data)
		return s.Width(maxLen).Align(lipgloss.Center)
	},
}

var _SUPPORTED_GLOBAL_MODIFIERS = maps.Keys(_GLOBAL_RENDER_MODIFIERS)

// Alter render function to apply global modifiers
func withGlobalModifiers(renderFn renderRowFunction) renderRowFunction {
	// return wrapped renderRowFunction
	return func(fields []interface{}, config renderRowConfig) ([]string, error) {
		global, local := findRecognizedModifiers(config.Modifiers)
		config.Modifiers = local
		rendered, err := renderFn(fields, config)
		if err != nil {
			return rendered, err
		}
		// build global style modifier
		style := lipgloss.NewStyle()
		for _, mod := range global {
			style = _GLOBAL_RENDER_MODIFIERS[mod](style, rendered)
		}
		// apply global modifier
		styled := make([]string, len(rendered))
		for i, r := range rendered {
			styled[i] = style.Render(r)
		}
		return styled, nil
	}
}

// Split into 'found' and 'not found' based on map
func findRecognizedModifiers(
	modifiers []string,
) ([]string, []string) {
	var found []string
	var missing []string
	for _, mod := range modifiers {
		if _, ok := _GLOBAL_RENDER_MODIFIERS[mod]; ok {
			found = append(found, mod)
		} else {
			missing = append(missing, mod)
		}
	}
	return found, missing
}

// Render modifier utility: get max len
func getMaxLen(data []string) int {
	maxLen := 0
	for _, d := range data {
		l := len(stripansi.Strip(d))
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}
