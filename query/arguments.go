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
