package filter

import "fmt"

/// Generic Value union type
/// MAKE SURE TO ACTUALLY DO UNION

type value interface {
	value()
	String() string
}

// String
type stringValue struct {
	Value string `@String`
}

func (v stringValue) value() {}
func (v stringValue) String() string {
	return v.Value
}

// Number
type numberValue struct {
	Value float64 `@(Float|Int) (?!"/")`
}

func (v numberValue) value() {}
func (v numberValue) String() string {
	return fmt.Sprintf("%f", v.Value)
}

// Boolean
type capturedBoolean bool

func (b *capturedBoolean) Capture(values []string) error {
	*b = values[0] == "TRUE"
	return nil
}

type booleanValue struct {
	Value capturedBoolean `@(TRUE|FALSE)`
}

func (v booleanValue) value() {}
func (v booleanValue) String() string {
	if v.Value {
		return "TRUE"
	}
	return "FALSE"
}

// Date
type dateValue struct {
	Year  int `@Int "/"`
	Month int `@Int "/"`
	Day   int `@Int`
}

func (v dateValue) value() {}
func (v dateValue) String() string {
	return fmt.Sprintf("%d/%d/%d", v.Year, v.Month, v.Day)
}

// RelativeDate
type relativeDateValue struct {
	Direction string `@("NEXT"|"LAST"|"PAST")`
	Amount    int    `@Int?`
	Unit      string `@("DAY"|"WEEK"|"MONTH"|"YEAR")`
}

func (r relativeDateValue) value() {}
func (r relativeDateValue) String() string {
	return fmt.Sprintf("%s %d %s", r.Direction, r.Amount, r.Unit)
}

// Empty/Null
type emptyValue struct {
	Value bool `@("EMPTY"|"NONE")`
}

func (v emptyValue) value() {}
func (v emptyValue) String() string {
	return "EMPTY"
}
