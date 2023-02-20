package query

import "flag"

/// FieldNames and FieldNamesConfig are LEGACY and will be removed

type FieldNames struct {
	TitleField    string
	DoneField     string
	DateField     string
	CategoryField string
	TagField      string
}

var FieldNamesConfig FieldNames

func SetupFieldNameFlags(flagsets []*flag.FlagSet) func() {
	names := FieldNames{}
	for _, fs := range flagsets {
		fs.StringVar(&names.TitleField, "title", "Name", "Name of title [text] field")
		fs.StringVar(&names.DoneField, "done", "Done", "Name of done [checkbox] field")
		fs.StringVar(&names.DateField, "date", "Due", "Name of the date [date] field")
		fs.StringVar(&names.CategoryField, "category", "Class", "Name of the category [select] field")
		fs.StringVar(&names.TagField, "tag", "Tags", "Name of the tag [multi-select] field")
	}
	return func() {
		FieldNamesConfig = names
	}
}

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
