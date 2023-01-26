package main

import (
	"flag"
	"strings"
)

/// Field Names Configuration: sets up how to parse task fields

var FieldNamesConfig FieldNames

func setupGlobalFieldNameFlags(flagsets []*flag.FlagSet) func() {
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

/// Tag Importance Configuration: sets up how to parse task tags

var ImportanceTags = ImportanceTagsMap{
	"exam":         HI,
	"projecttask":  HI,
	"presentation": HI,
	"project":      HI,
	"paper":        HI,
	"meeting":      LO,
	"read":         LO,
	"utility":      LO,
}

func setupGlobalTagImportanceFlags(flagsets []*flag.FlagSet) func() {
	importantTags := ""
	unimportantTags := ""

	for _, fs := range flagsets {
		fs.StringVar(&importantTags, "important", "exam,projecttask,presentation,project", "Comma-separated tag names to render as important")
		fs.StringVar(&unimportantTags, "unimportant", "meeting,read,utility", "Comma-separated tag names to render as unimportant")
	}

	return func() {
		for _, tag := range strings.Split(importantTags, ",") {
			ImportanceTags[tag] = HI
		}
		for _, tag := range strings.Split(unimportantTags, ",") {
			ImportanceTags[tag] = LO
		}
	}
}
