package query

import (
	"fmt"
	"strings"
)

func formatSortDirective(sortString SortString) (string, error) {
	// from "Due:desc,Name"	=> '[{"property": "Due", "direction": "descending"}, {"property": "Name", "direction": "ascending"}]'
	var sortDirective string = "["
	for i, prop := range strings.Split(sortString, ",") {
		sp := strings.Split(prop, ":")
		field, dir := sp[0], "ascending"
		if len(sp) > 1 {
			switch sp[1] {
			case "desc":
				dir = "descending"
			case "asc":
				dir = "ascending"
			default:
				return "", fmt.Errorf(errType.InvalidSortDirection, sp[1])
			}
		}
		if len(sp) > 2 {
			return "", fmt.Errorf(errType.InvalidSortSyntax, prop)
		}
		if i != 0 {
			sortDirective += ", "
		}
		sortDirective += fmt.Sprintf(`{"property": "%s", "direction": "%s"}`, field, dir)
	}
	sortDirective += "]"
	return sortDirective, nil
}
