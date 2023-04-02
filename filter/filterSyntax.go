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
	// note validity is checked in validate.go
	Type string `@("select"|"checkbox"|"number"|"text"|"date"|"multiselect"|Ident)`
}

type fieldTypeString string

const (
	SelectType      = fieldTypeString("select")
	CheckboxType    = fieldTypeString("checkbox")
	NumberType      = fieldTypeString("number")
	TextType        = fieldTypeString("text")
	DateType        = fieldTypeString("date")
	MultiselectType = fieldTypeString("multiselect")
)

// Filter Operator

type operator struct {
	Not bool   `@("NOT"|"!")?`
	Op  string `@("="|("<" "="?)|(">" "="?)|"CONTAINS"|"STARTS_WITH"|"ENDS_WITH"|Ident)`
}

func (o *operator) String() string {
	if o.Not {
		return "!" + o.Op
	}
	return o.Op
}

// Filter Value: defined in valueSyntax.go

/// FILTER RENDER

func (f *filter) Render() (string, error) {
	// get easy base values
	propertyName := string(*f.Field)
	typeName := fieldTypeString(f.Type.Type)
	// check if valid primitives
	if err := ensureValidOperator(f.Operator); err != nil {
		return "", err
	}
	if err := ensureValidType(typeName); err != nil {
		return "", err
	}
	// build filter condition
	condition, err := getFilterCondition(f, typeName)
	if err != nil {
		return "", err
	}
	if t, ok := typeNameOverride[f.Type.Type]; ok {
		typeName = fieldTypeString(t)
	}
	return fmt.Sprintf(template, propertyName, typeName, condition), nil
}

const template = `{"property": "%s", "%s": {%s}}`

var typeNameOverride = map[string]string{
	"multiselect": "multi-select",
	"text":        "title", // TODO: could also be rich_text?
}

/// DEFINITIONS for assigning notion filters to syntax

func getFilterCondition(f *filter, t fieldTypeString) (string, error) {
	// check "operation value" overrides -- like '= EMPTY'
	if found, ok := operationValue[f.Operator.String()+" "+f.Value.String()]; ok {
		if err := ensureValidTypeValue(t, f.Value); err != nil {
			return "", err
		}
		return found, nil
	}
	// check unique op keywords
	if specialOp, ok := typeOperationKeyword[t][f.Operator.String()]; ok {
		return fmt.Sprintf(`"%s": %s`, specialOp, f.Value.Render()), nil
	}
	// check default op keywords
	if defaultOp, ok := defaultOperationKeyword[f.Operator.String()]; ok {
		if err := ensureValidTypeOperator(t, f.Operator); err != nil {
			return "", err
		}
		if err := ensureValidTypeValue(t, f.Value); err != nil {
			return "", err
		}
		return fmt.Sprintf(`"%s": %s`, defaultOp, f.Value.Render()), nil
	}
	// bad operation... fail
	return "", fmt.Errorf(errType.InvalidSyntax, t)
}

var operationValue = map[string]string{
	"= EMPTY":  `"is_empty": true`,
	"!= EMPTY": `"is_not_empty": true`,
}

var typeOperationKeyword = map[fieldTypeString]map[string]string{
	DateType: {
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
