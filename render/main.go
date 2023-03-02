package render

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mfigurski80/NTPeek/priority"
	"github.com/mfigurski80/NTPeek/types"
	"github.com/muesli/termenv"
	"golang.org/x/exp/maps"
)

/// RenderTasks renders a list of tasks

var selectRenderRegex = regexp.MustCompile(`%([a-zA-Z0-9:]+)%`)

func RenderTasks(tasks []types.NotionEntry, selectRender SelectRenderString, priorityConfig priority.PriorityConfig) {
	lipgloss.SetColorProfile(termenv.TrueColor)
	// find fields needed for render
	m := selectRenderRegex.FindAllString(string(selectRender), -1)
	fields := make([]string, len(m))
	for i, v := range m {
		fields[i] = v[1 : len(v)-1]
	}
	// render field data
	renderedFields := getRenderedFields(tasks, fields, priorityConfig)
	// place field data into render string
	formatString := selectRenderRegex.ReplaceAllString(string(selectRender), "%s")
	ret := sprintfList(formatString+"\n", renderedFields)
	fmt.Printf(ret)
}

func sprintfList(format string, list [][]string) string {
	result := ""
	for i := range list[0] {
		row := make([]interface{}, len(list))
		for j, column := range list {
			row[j] = column[i]
		}
		result += fmt.Sprintf(format, row...)
	}
	return result
}

func getRenderedFields(tasks []types.NotionEntry, fields []string, priorityConfig priority.PriorityConfig) [][]string {
	if len(tasks) == 0 {
		return make([][]string, len(fields))
	}
	// get priorities
	priorities := priority.Assign(tasks, priorityConfig)
	// parse each field: NAME[:MODIFIER]*
	fieldNames := make([]string, len(fields))
	fieldModifiers := make([][]string, len(fields))
	for i, field := range fields {
		p := strings.Split(field, ":")
		fieldNames[i], fieldModifiers[i] = p[0], p[1:]
	}
	// get values list for each interesting field
	fieldVals := make([][]interface{}, len(fields))
	for i, name := range fieldNames {
		fieldVals[i] = make([]interface{}, len(tasks))
		if tasks[0][name] == nil {
			fmt.Printf(errStart+err.FieldNotFound, name, maps.Keys(tasks[0]))
			continue
		}
		for j, task := range tasks {
			fieldVals[i][j] = task[name]
		}
	}
	// render each field
	rendered := make([][]string, len(fields))
	for i, name := range fieldNames {
		renderFunc, err := getGenericRenderFunc(fieldVals[i], name)
		if err != nil {
			fmt.Printf(err.Error())
			rendered[i] = make([]string, len(fieldVals[i]))
			continue
		}
		rendered[i], err = renderFunc(fieldVals[i], renderRowConfig{
			name, fieldModifiers[i], priorities,
		})
		if err != nil {
			fmt.Printf(err.Error())
		}
	}
	// return rendered fields
	return rendered
}
