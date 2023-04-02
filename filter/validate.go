package filter

import (
	"fmt"

	"golang.org/x/exp/maps"
)

/// BASIC CHECKS -- for primitives

func ensureValidOperator(op *operator) error {
	if _, ok := defaultOperationKeyword[op.String()]; !ok {
		if _, ok := negatableOperators[op.Op]; op.Not && !ok {
			return fmt.Errorf(errType.NonNegatableOp, op.Op)
		}
		return fmt.Errorf(errType.InvalidKeyword, "Operator", op.Op, maps.Keys(defaultOperationKeyword))
	}
	return nil
}

var negatableOperators = map[string]struct{}{
	"=":        {},
	"CONTAINS": {},
}

func ensureValidType(t fieldTypeString) error {
	if _, ok := supportedTypeOperations[t]; !ok {
		return fmt.Errorf(errType.InvalidKeyword, "Type", t, maps.Keys(supportedTypeOperations))
	}
	return nil
}

/// TYPE/OP CHECKS

func ensureValidTypeOperator(t fieldTypeString, op *operator) error {
	if _, ok := supportedTypeOperations[t][op.String()]; !ok {
		k := maps.Keys(supportedTypeOperations[t])
		// check special cases of type operations
		if ops, ok := typeOperationKeyword[t]; ok {
			k = append(k, maps.Keys(ops)...)
		}
		return fmt.Errorf(errType.TypeOpMismatch, t, op.String(), k)
	}
	return nil
}

var supportedTypeOperations = map[fieldTypeString]map[string]struct{}{
	DateType:        {"=": {}}, // additional ops over typeOperationKeyword
	CheckboxType:    {"=": {}, "!=": {}},
	SelectType:      {"=": {}, "!=": {}},
	NumberType:      {"=": {}, "!=": {}, ">": {}, ">=": {}, "<": {}, "<=": {}},
	MultiselectType: {"CONTAINS": {}, "!CONTAINS": {}},
	TextType:        {"=": {}, "!=": {}, "CONTAINS": {}, "!CONTAINS": {}, "STARTS_WITH": {}, "ENDS_WITH": {}},
}

/// TYPE/VAL CHECKS

func ensureValidTypeValue(t fieldTypeString, val value) error {
	// check val whitelist
	switch val.(type) {
	case emptyValue:
		if t == CheckboxType {
			return fmt.Errorf(errType.TypeValMismatch, t, val)
		}
		return nil
	}
	// check type whitelist
	ok := false
	switch t {
	case TextType:
		_, ok = val.(stringValue)
	case NumberType:
		_, ok = val.(numberValue)
	case CheckboxType:
		_, ok = val.(booleanValue)
	case DateType:
		_, ok1 := val.(dateValue)
		_, ok2 := val.(relativeDateValue)
		ok = ok1 || ok2
	case SelectType:
		_, ok = val.(stringValue)
	case MultiselectType:
		_, ok = val.(stringValue)
	}
	if ok {
		return nil
	}
	return fmt.Errorf(errType.TypeValMismatch, t, val)
}
