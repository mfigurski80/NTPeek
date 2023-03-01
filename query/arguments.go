package query

import "flag"

/// `--sort` flag provides way to sort by multiple fields

// Represents format like: "Due:asc,Name"
type SortString = string

func SetupSortFlag(flagsets []*flag.FlagSet) *SortString {
	var sort string
	for _, fs := range flagsets {
		fs.StringVar(&sort, "sort", "Due:asc,Name", "Comma-separate list of fields to sort by, with optional order (asc/desc)")
	}
	return &sort
}

/// `--filter` flag provides way to filter by multiple fields

// Represents grammar string defined in query/filter/*Syntax.go
type FilterString = string

func SetupFilterFlag(flagsets []*flag.FlagSet) *FilterString {
	var filter string
	for _, fs := range flagsets {
		fs.StringVar(&filter, "filter", "Done:checkbox = FALSE AND Due:date < NEXT 10 DAY", "Filter query language: see github for full documentation")
	}
	return &filter
}
