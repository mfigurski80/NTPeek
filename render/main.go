package render

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/mfigurski80/NTPeek/types"
	"github.com/muesli/termenv"
)

type NotionEntry = types.NotionEntry

/// RenderTasks renders a list of tasks

func RenderTasks(tasks []types.NotionEntry, selectRender SelectRenderString) {
	lipgloss.SetColorProfile(termenv.TrueColor)
	// build map of names for each field
	fieldNames := make([]string, len(tasks[0]))
	fieldVals := make([][]interface{}, len(tasks[0]))
	i := 0
	for k, _ := range tasks[0] {
		fieldNames[i] = k
		fieldVals[i] = make([]interface{}, len(tasks))
		for j, task := range tasks {
			fieldVals[i][j] = task[k]
		}
		i++
	}
	// TODO: filter out fields we need to render
	// render each column into strings
	vals := make([][]string, len(fieldNames))
	for i, name := range fieldNames {
		if name == "_id" {
			continue
		}
		vals[i] = make([]string, len(tasks))
		rawVals := fieldVals[i].([]map[string]interface{})
		fmt.Printf("%s - %s\n", name, rawVals["type"])
		fmt.Println(fieldVals[i][0])
		renderFunc := getFieldRenderFunc(rawVals)
	}
	return
	//
	// for _, task := range tasks {
	// selectRender.renderTask(&task, priority.LO)
	// return
	// }
}

type RenderRowFunction func([]interface{}, []string) []string

func getFieldRenderFunc(field []map[string]interface{}) RenderRowFunction {
	switch field[0]["type"].(string) {
	case "title":
		fmt.Println("title")
		return renderTitle
	case "rich_text":
		fmt.Println("rich_text")
		return renderTitle
	case "select":
		fmt.Println("select")
		return renderSelect
	// case "multi_select":
	// fmt.Println("multi_select")
	// return []string{}
	case "date":
		fmt.Println("date")
		return renderDate
	default:
		fmt.Println("unknown")
		return func([]interface{}, []string) []string { return []string{} }
	}
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
