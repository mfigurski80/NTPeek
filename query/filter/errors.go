package filter

var errType = struct {
	InvalidSyntax   string
	InvalidKeyword  string
	TypeValMismatch string
	TypeOpMismatch  string
}{
	InvalidSyntax:   "ERROR creating filter: invalid syntax: '%s'",
	InvalidKeyword:  "ERROR creating filter: invalid %s keyword '%s'",
	TypeValMismatch: "ERROR creating filter: type/value mismatch for '%s' and '%s'",
	TypeOpMismatch:  "ERROR creating filter: type/operator mismatch for '%s' and '%s'. Type supports: %v",
}
