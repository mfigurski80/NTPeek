package filter

import (
	"flag"
	"fmt"
	"strings"
)

/// `--filter` flag provides way to filter by multiple fields

// Represents grammar string defined in query/filter/*Syntax.go
type FilterString = string

type FilterFlags []FilterString

func (f *FilterFlags) String() string {
	return fmt.Sprintf(`"%s"`, strings.Join(*f, ","))
}

var defaultValue = FilterFlags{"Done:checkbox = FALSE AND Due:date < NEXT 10 DAY"}
var replacedDefaultFilterValue = false

func (f *FilterFlags) Set(value string) error {
	if !replacedDefaultFilterValue {
		*f = FilterFlags{}
		replacedDefaultFilterValue = true
	}
	*f = append(*f, value)
	return nil
}

func SetupFilterFlag(flagsets []*flag.FlagSet) *FilterFlags {
	var filters FilterFlags = defaultValue
	for _, fs := range flagsets {
		fs.Var(&filters, "filter", "Filters following query language: see github for full documentation")
	}
	return &filters
}
