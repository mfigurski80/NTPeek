package filter

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
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
	Not bool   `@("NOT"|"!")?`
	Op  string `@("="|("<" "="?)|(">" "="?)|"CONTAINS"|"STARTS_WITH"|"ENDS_WITH")`
}

func (o *operator) String() string {
	if o.Not {
		return "!" + o.Op
	}
	return o.Op
}

// Filter Value: defined in valueSyntax.go

var negatableOperators = map[string]struct{}{
	"=":        {},
	"CONTAINS": {},
}

/// FILTER RENDER

func (f *filter) Render() (string, error) {
	template := `{"property": "%s", "%s": {%s}}`
	// get property name
	propertyName := string(*f.Field)
	// get notion typename
	typeName := fieldTypeString(f.Type.Type)
	ntTypeName := typeName
	if t, ok := typeNameOverride[f.Type.Type]; ok {
		ntTypeName = fieldTypeString(t)
	}
	// check if valid operator
	if f.Operator.Not {
		if _, ok := negatableOperators[f.Operator.Op]; !ok {
			return "", fmt.Errorf(errType.NonNegatableOp, f.Operator.Op)
		}
	}
	// get condition: in-order check value, type, op
	condition := ""
	if c, ok := operationValue[f.Operator.String()+" "+f.Value.String()]; ok {
		// op/value check: some (like = EMPTY) correspond to unique filter
		condition = c
	} else if c, ok := typeOperationKeyword[typeName][f.Operator.String()]; ok {
		// type/op check: some (like date =) correspond to unique filter
		condition = fmt.Sprintf(`"%s": %s`, c, f.Value.Render())
	} else if c, ok := defaultOperationKeyword[f.Operator.String()]; ok {
		// default: ensure type supports op, type matches value
		if _, ok := supportedTypeOperations[typeName][f.Operator.String()]; !ok {
			// if not, list supported ops
			k := maps.Keys(supportedTypeOperations[typeName])
			if ops, ok := typeOperationKeyword[typeName]; ok {
				k = append(k, maps.Keys(ops)...)
			}
			return "", fmt.Errorf(errType.TypeOpMismatch, typeName, f.Operator.String(), k)
		}
		if matches := valueMatchesType(f.Value, typeName); !matches {
			return "", fmt.Errorf(errType.TypeValMismatch, typeName, f.Value)
		}
		condition = fmt.Sprintf(`"%s": %s`, c, f.Value.Render())
	} else {
		return "", fmt.Errorf(errType.InvalidKeyword, "Operator", f.Operator.String())
	}
	return fmt.Sprintf(template, propertyName, ntTypeName, condition), nil
}

var typeNameOverride = map[string]string{
	"multiselect": "multi-select",
	"text":        "title", // TODO: could also be rich_text?
}

/// DEFINITIONS for assigning notion filters to syntax

var operationValue = map[string]string{
	"= EMPTY":  `"is_empty": true`,
	"!= EMPTY": `"is_not_empty": true`,
}

var typeOperationKeyword = map[fieldTypeString]map[string]string{
	Date: {
		">":  "after",
		">=": "on_or_after",
		"<":  "before",
		"<=": "on_or_before",
	},
}

// given need default operation keyword, ensure type supports it
var supportedTypeOperations = map[fieldTypeString]map[string]struct{}{
	Date:        {"=": {}}, // additional ops over typeOperationKeyword
	Checkbox:    {"=": {}, "!=": {}},
	Select:      {"=": {}, "!=": {}},
	Number:      {"=": {}, "!=": {}, ">": {}, ">=": {}, "<": {}, "<=": {}},
	Multiselect: {"CONTAINS": {}, "!CONTAINS": {}},
	Text:        {"=": {}, "!=": {}, "CONTAINS": {}, "!CONTAINS": {}, "STARTS_WITH": {}, "ENDS_WITH": {}},
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
