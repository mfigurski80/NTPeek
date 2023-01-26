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
var DefaultImportance = LO

func setupGlobalTagImportanceFlags(flagsets []*flag.FlagSet) func() {
	importantTags := ""
	unimportantTags := ""
	dImportance := true

	for _, fs := range flagsets {
		fs.StringVar(&importantTags, "important", "", "Comma-separated tags to render as important (Default \"exam,projecttask,presentation,project,paper)\"")
		fs.StringVar(&unimportantTags, "unimportant", "", "Comma-separated tags to render as unimportant (Default \"meeting,read,utility\")")
		fs.BoolVar(&dImportance, "default-importance", false, "Default avg importance if no tags are present")
	}

	return func() {
		if importantTags != "" {
			// remove all imporant tags
			for k := range ImportanceTags {
				if ImportanceTags[k] == HI {
					delete(ImportanceTags, k)
				}
			}
			for _, tag := range strings.Split(importantTags, ",") {
				ImportanceTags[tag] = HI
			}
		}
		if unimportantTags != "" {
			// remove all unimportant tags
			for k := range ImportanceTags {
				if ImportanceTags[k] == LO {
					delete(ImportanceTags, k)
				}
			}
			for _, tag := range strings.Split(unimportantTags, ",") {
				ImportanceTags[tag] = LO
			}
		}
		if dImportance {
			DefaultImportance = AVG
		}
	}
}
