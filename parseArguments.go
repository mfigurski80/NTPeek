package main

import (
	"flag"
)

type FieldNames struct {
	TitleField    string
	DoneField     string
	DateField     string
	CategoryField string
	TagField      string
}

func parseFieldNameArguments(args []string) FieldNames {
	titleField := flag.String(
		"titleField", "Name", "title [text] field of the task")
	doneField := flag.String(
		"doneField", "Done", "done [checkbox] field of the task")
	dateField := flag.String(
		"dateField", "Due", "date [date] field of the task")
	categoryField := flag.String(
		"categoryField", "Class", "category [select] field of the task")
	tagField := flag.String(
		"tagField", "Tags", "tag [multi-select] field of the task")
	flag.Parse()
	return FieldNames{
		*titleField, *doneField, *dateField, *categoryField, *tagField,
	}
}
