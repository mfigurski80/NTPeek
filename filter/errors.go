package filter

var errType = struct {
	InvalidSyntax   string
	InvalidKeyword  string
	TypeValMismatch string
	TypeOpMismatch  string
	NonNegatableOp  string
}{
	InvalidSyntax:   "ERROR creating filter: invalid syntax: '%s'",
	InvalidKeyword:  "ERROR creating filter: invalid '%s' keyword '%s'. Tool supports only: %v",
	TypeValMismatch: "ERROR creating filter: type/value mismatch for '%s' and '%s'",
	TypeOpMismatch:  "ERROR creating filter: type/operator mismatch for '%s' and '%s'. Type supports: %v",
	NonNegatableOp:  "ERROR creating filter: operator '%s' cannot be negated with neither ! or NOT",
}
