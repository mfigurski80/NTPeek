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

/// `--limit` flag lets users limit number of items displayed

type LimitNumber = uint

func SetupLimitFlag(flagsets []*flag.FlagSet) *LimitNumber {
	var limit uint
	for _, fs := range flagsets {
		fs.UintVar(&limit, "limit", 100, "Limit number of items displayed")
	}
	return &limit
}
