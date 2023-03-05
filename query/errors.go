package query

var errType = struct {
	InvalidSortDirection string
	InvalidSortSyntax    string
}{
	InvalidSortDirection: "ERROR formatting sort: invalid direction '%s'. Use 'desc' or 'asc'",
	InvalidSortSyntax:    "ERROR formatting sort: invalid syntax '%s'. Use 'field:direction, field:direction'",
}
