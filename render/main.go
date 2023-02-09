package render

import (
	"fmt"
	"regexp"

	"github.com/charmbracelet/lipgloss"
	"github.com/mfigurski80/NTPeek/types"
	"github.com/muesli/termenv"
)

/// RenderTasks renders a list of tasks

var selectRenderRegex = regexp.MustCompile(`%([a-zA-Z0-9.]+)%`)

func RenderTasks(tasks []types.NotionEntry, selectRender SelectRenderString) {
	lipgloss.SetColorProfile(termenv.TrueColor)
	// find fields needed for render
	m := selectRenderRegex.FindAllString(string(selectRender), -1)
	fields := make([]string, len(m))
	for i, v := range m {
		fields[i] = v[1 : len(v)-1]
	}
	// render field data
	renderedFields := getRenderedFields(tasks, fields)
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

func zip(a ...[]string) [][]string {
	var zipped [][]string
	for i := 0; i < len(a[0]); i++ {
		var row []string
		for _, v := range a {
			row = append(row, v[i])
		}
		zipped = append(zipped, row)
	}
	return zipped
}

/// Replaces every match with a field-specific render of its value

// var selectRenderRegex = regexp.MustCompile(`%([a-zA-Z0-9.]+)%`)
//
// func (s SelectRenderString) renderTask(t *NotionEntry, importance priority.Priority) string {
// for k := range *t {
// fmt.Print(k)
// }
// fmt.Println()
// m := selectRenderRegex.FindAllString(string(s), -1)
// for _, v := range m {
// // split on . to get class and method
// split := strings.Split(v[1:len(v)-1], ".")
// fmt.Printf("%s :: %v\n", split[0], (*t)[split[0]])
//
// // fmt.Println((*t)[split[0]])
// }
// return ""
// }

// func printTasks(tasks []NotionEntry, selectRender SelectRenderString) {
//
// maxClassLen := 0
// classLengths := make([]int, len(tasks))
// for i, task := range tasks {
// if len(task.Class) > maxClassLen {
// maxClassLen = len(task.Class)
// }
// classLengths[i] = len(task.Class)
// }
// for _, task := range tasks {
// // GET CLASS + FORMAT
// hi := colorMap[task.ClassColor]
// class := lipgloss.NewStyle().
// Background(lipgloss.Color(hi.Bg)).
// Foreground(lipgloss.Color(hi.Fore)).
// Render(task.Class)
// class = lipgloss.NewStyle().
// Width(maxClassLen).
// Align(lipgloss.Right).
// Render(class)
//
// // GET IMPORTANCE
// importanceVal := parseImportance(task)
// importance := formatImportance(
// importanceVal,
// [3]string{"│ ", "│ ", "│!"},
// )
//
// // GET TASK ID
// // id := lipgloss.NewStyle().
// // Faint(true).
// // Render(fmt.Sprintf("%.2s", task.Id))
//
// // PRINT
// // fmt.Printf("%s %s %s\n", class, name, due)
// fmt.Printf("%s %s%s  %s\n", class, importance, name, due)
//
// }
// }
