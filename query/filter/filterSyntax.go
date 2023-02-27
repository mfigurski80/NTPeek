package filter

import (
	"fmt"
	"strings"
)

type filter struct {
	Field    *fieldName `@(Ident)+ ":"`
	Type     *fieldType `@@`
	Operator *operator  `@@`
	Value    value      `@@`
}

func (f *filter) String() string {
	return f.Type.Type + "(" + string(*f.Field) + ") " + f.Operator.Op + " " + f.Value.String()
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

type fieldTypeString string

const (
	Select      = fieldTypeString("select")
	Checkbox    = fieldTypeString("checkbox")
	Number      = fieldTypeString("number")
	Text        = fieldTypeString("text")
	Date        = fieldTypeString("date")
	Multiselect = fieldTypeString("multiselect")
)

// Filter Operator

type operator struct {
	Not bool   `@("NOT"|"!" (?= "="|"CONTAINS"))?`
	Op  string `@("="|("<" "="?)|(">" "="?)|"CONTAINS"|"STARTS_WITH"|"ENDS_WITH")`
}

func (o *operator) String() string {
	if o.Not {
		return "!" + o.Op
	}
	return o.Op
}

// Filter Value: defined in valueSyntax.go

/// FILTER RENDER

func (f *filter) Render() string {
	template := `{"property": "%s", "%s": {%s}}`
	// get property name
	propertyName := string(*f.Field)
	// get typename
	typeName := f.Type.Type
	if t, ok := typeNameOverride[f.Type.Type]; ok {
		typeName = t
	}
	// get condition: in-order check value, type, op
	condition := ""
	if c, ok := operationValue[f.Operator.String()+" "+f.Value.String()]; ok {
		condition = c
	} else if c, ok := typeOperationKeyword[typeName][f.Operator.String()]; ok {
		condition = fmt.Sprintf(`"%s": %s`, c, f.Value.Render())
	} else if c, ok := defaultOperationKeyword[f.Operator.String()]; ok {
		condition = fmt.Sprintf(`"%s": %s`, c, f.Value.Render())
	} else {
		panic(fmt.Sprintf("invalid filter: %s", f.String()))
	}
	// return render
	return fmt.Sprintf(template, propertyName, typeName, condition)
}

var typeNameOverride = map[string]string{
	"multiselect": "multi-select",
	"text":        "title", // TODO: could also be rich_text?
}

var operationValue = map[string]string{
	"= EMPTY":  `"is_empty": true`,
	"!= EMPTY": `"is_not_empty": true`,
}

var typeOperationKeyword = map[string]map[string]string{
	"date": {
		"!=": "",
		">":  "after",
		">=": "on_or_after",
		"<":  "before",
		"<=": "on_or_before",
	},
}

var defaultOperationKeyword = map[string]string{
	"=":           "equals",
	"!=":          "does_not_equal",
	">":           "greater_than",
	">=":          "greater_than_or_equal_to",
	"<":           "less_than",
	"<=":          "less_than_or_equal_to",
	"CONTAINS":    "contains",
	"!CONTAINS":   "does_not_contain",
	"STARTS_WITH": "starts_with",
	"ENDS_WITH":   "ends_with",
}
