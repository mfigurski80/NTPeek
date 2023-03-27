package filter

import (
	"fmt"
	"time"

	"github.com/alecthomas/participle/v2"
)

var timeProvider = time.Now

/// Generic Value union type
/// MAKE SURE TO ACTUALLY DO UNION IN PARSER

type value interface {
	Render() string
	String() string
}

var valueUnion = participle.Union[value](stringValue{}, numberValue{}, booleanValue{}, dateValue{}, relativeDateValue{}, emptyValue{})

// String
type stringValue struct {
	Value string `@String`
}

func (v stringValue) String() string {
	return v.Value
}
func (v stringValue) Render() string {
	return fmt.Sprintf(`"%s"`, v.Value)
}

// Number
type numberValue struct {
	Value float64 `@(Float|Int) (?!"/")`
}

func (v numberValue) String() string {
	return fmt.Sprintf("%f", v.Value)
}
func (v numberValue) Render() string {
	return fmt.Sprintf("%f", v.Value)
}

// Boolean
type capturedBoolean bool

func (b *capturedBoolean) Capture(values []string) error {
	*b = values[0] == "TRUE"
	return nil
}

type booleanValue struct {
	Value capturedBoolean `@("TRUE"|"FALSE")`
}

func (v booleanValue) String() string {
	if v.Value {
		return "TRUE"
	}
	return "FALSE"
}
func (v booleanValue) Render() string {
	if v.Value {
		return "true"
	}
	return "false"
}

// Date
type dateValue struct {
	Year  int `@Int "/"`
	Month int `@Int "/"`
	Day   int `@Int`
}

func (v dateValue) String() string {
	return fmt.Sprintf("%04d/%02d/%02d", v.Year, v.Month, v.Day)
}
func (v dateValue) Render() string {
	return fmt.Sprintf(`"%04d-%02d-%02d"`, v.Year, v.Month, v.Day)
}

// RelativeDate
type relativeDateValue struct {
	Direction string `@("NEXT"|"LAST"|"PAST")`
	Amount    int    `@Int?`
	Unit      string `@("DAY"|"WEEK"|"MONTH"|"YEAR")`
}

func (r relativeDateValue) String() string {
	return fmt.Sprintf("%s %d %s", r.Direction, r.Amount, r.Unit)
}
func (r relativeDateValue) Render() string {
	dir := -1
	if r.Direction == "NEXT" {
		dir = 1
	}
	if r.Amount == 0 {
		r.Amount = 1
	}
	t := timeProvider()
	switch r.Unit {
	case "DAY":
		t = t.AddDate(0, 0, dir*r.Amount)
	case "WEEK":
		t = t.AddDate(0, 0, dir*7*r.Amount)
	case "MONTH":
		t = t.AddDate(0, dir*r.Amount, 0)
	case "YEAR":
		t = t.AddDate(dir*r.Amount, 0, 0)
	}
	return t.Format(`"2006-01-02"`)
}

// Empty/Null
type emptyValue struct {
	Value bool `@("EMPTY"|"NONE")`
}

func (v emptyValue) String() string {
	return "EMPTY"
}
func (v emptyValue) Render() string {
	return "EMPTY_VALUE_RENDER_NOT_USED" // wont be used
}

/// Type/Value validation

func valueMatchesType(val value, typ fieldTypeString) bool {
	switch typ {
	case Text:
		_, ok := val.(stringValue)
		return ok
	case Number:
		_, ok := val.(numberValue)
		return ok
	case Checkbox:
		_, ok := val.(booleanValue)
		return ok
	case Date:
		_, ok1 := val.(dateValue)
		_, ok2 := val.(relativeDateValue)
		return ok1 || ok2
	case Select:
		_, ok := val.(stringValue)
		return ok
	case Multiselect:
		_, ok := val.(stringValue)
		return ok
	default:
		return false
	}
}
