package query

import (
	"github.com/mfigurski80/NTPeek/query/filter"
)

func formatFilterDirective(filterString string) string {
	// from "Due:date < 2022/02/02 AND Done:checkbox = FALSE" to
	// `{ "and": [
	// 		{"property": "Due", "date": {"before": "2022/02/02"},
	//		{"property": "Done": "checkbox": { "equals": false }}
	// ] }`
	f := filter.ParseFilter(filterString)
	// fmt.Println(f)
	return f
}
