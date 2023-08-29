package renderField

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var loc, _ = time.LoadLocation("Local")

var overdueDateStyle = map[bool]lipgloss.Style{
	true: lipgloss.NewStyle().
		Background(lipgloss.Color(colorMap["red"].Bg)).
		Foreground(lipgloss.Color(colorMap["red"].Fg)),
	false: lipgloss.NewStyle().
		Faint(true),
}

/// Render date field relative to local time

func RenderDate(fields []interface{}, config RenderRowConfig) ([]string, error) {
	res := make([]string, len(fields))
	// parse modifiers
	stringifyStrategy := _RELATIVE
	for _, mod := range config.Modifiers {
		switch mod {
		case "relative", "rel":
			stringifyStrategy = _RELATIVE
		case "simple", "sim":
			stringifyStrategy = _SIMPLE
		case "full", "ful":
			stringifyStrategy = _FULL
		default:
			return res, fmt.Errorf(
				errType.UnsupportedMod, config.Name, "date", mod,
				append(_SUPPORTED_GLOBAL_MODIFIERS, "relative", "simple", "full"),
			)
		}
	}
	// render into result
	var gErr error
	for i, field := range fields {
		f, ok := field.(map[string]interface{})["date"].(map[string]interface{})
		if !ok {
			res[i] = ""
			gErr = fmt.Errorf(errType.Internal, config.Name, field)
			continue
		}
		date := f["start"].(string)
		t, err := time.ParseInLocation("2006-01-02", date[:10], loc)
		if err != nil {
			return res, fmt.Errorf(errType.Internal, config.Name, err.Error())
		}
		relative := fmt.Sprintf("(%s)", stringifyDateMap[stringifyStrategy](t))
		res[i] = overdueDateStyle[isOverdue(t)].Render(relative)
	}
	return res, gErr
}

/// Different ways of stringifying date

type stringifyDateOption uint8

const (
	_RELATIVE stringifyDateOption = iota
	_SIMPLE
	_FULL
)

var stringifyDateMap = map[stringifyDateOption](func(time.Time) string){
	_RELATIVE: func(t time.Time) string {
		hoursDiff := time.Until(t).Hours()
		diff := math.Ceil((hoursDiff) / 24)
		if diff == 0 {
			return "Today"
		}
		if diff == 1 {
			return "Tomorrow"
		}
		if diff == -1 {
			return "Yesterday"
		}
		if diff < 10 && diff > -7 {
			if diff >= 8 {
				return "next " + t.Weekday().String()
			}
			if diff > 0 {
				return t.Weekday().String()
			}
			return "last " + t.Weekday().String()
		}
		if diff > 0 {
			return "in " + fmt.Sprint(diff) + " days"
		}
		return fmt.Sprint(-diff) + " days ago"
	},
	_SIMPLE: func(t time.Time) string {
		return t.Format("Jan 2")
	},
	_FULL: func(t time.Time) string {
		return t.Format("2006-01-02")
	},
}

// overdue utility
func isOverdue(t time.Time) bool {
	// Everything due at midnight: only overdue if due yesterday
	return time.Until(t).Hours() < -24
}
