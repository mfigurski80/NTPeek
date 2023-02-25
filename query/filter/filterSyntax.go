package filter

import "strings"

type Filter struct {
	Field    *fieldName `@(Ident)+ ":"`
	Type     *fieldType `@@`
	Operator *operator  `@@`
	Value    value      `@@`
}

// Field Name

type fieldName string

func (f *fieldName) Capture(values []string) error {
	if len(*f) > 0 {
		*f += " "
	}
	*f += fieldName(strings.Join(values, " "))
	return nil
}

// Field Type

type fieldType struct {
	Type string `@("select"|"checkbox"|"number"|"text"|"date"|"multiselect")`
}

// Filter Operator

type operator struct {
	Not bool   `@("NOT"|"!" (?= "="|"CONTAINS"))?`
	Op  string `@("="|("<" "="?)|(">" "="?)|"CONTAINS"|"STARTS_WITH"|"ENDS_WITH")`
}

// Filter Value: defined in valueSyntax.go
