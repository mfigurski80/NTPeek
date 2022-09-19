package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func parseNotionRichText(richText []interface{}) string {
	var text string
	for _, entry := range richText {
		text += entry.(map[string]interface{})["plain_text"].(string)
	}
	return text
}

type Task struct {
	Name       string
	Due        string
	Class      string
	ClassColor string
}

func queryNotionTaskDB() []Task {
	url := "https://api.notion.com/v1/databases/d048f752003e4c199533c9a39608917e/query"
	payload := strings.NewReader(`{
		"page_size": 100,
		"filter": {
			"and": [{
				"property": "Done",
				"checkbox": {
					"equals": false
				}
			}, {
				"property": "Due",
				"date": {
					"next_week": {}
				}
			}]
		},
		"sorts": [{
			"property": "Due",
			"direction": "ascending"
		}]
	}`)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	if result["object"] == "error" {
		fmt.Printf("Returned Error [%v]: %s\n", result["status"], result["message"])
	}

	var tasks []Task
	for _, entry := range result["results"].([]interface{}) {
		properties := entry.(map[string]interface{})["properties"].(map[string]interface{})
		name := parseNotionRichText(properties["Name"].(map[string]interface{})["title"].([]interface{}))
		due := properties["Due"].(map[string]interface{})["date"].(map[string]interface{})["start"].(string)
		class := properties["Class"].(map[string]interface{})["select"].(map[string]interface{})["name"].(string)
		classColor := properties["Class"].(map[string]interface{})["select"].(map[string]interface{})["color"].(string)
		// fmt.Printf("(%s) %s due %s\n", class, name, due)
		tasks = append(tasks, Task{name, due, class, classColor})
	}
	return tasks
}

func formatRelativeDate(date string) string {
	// convert "2022-07-01" to "next Monday"
	t, _ := time.Parse("2006-01-02", date)
	diff := t.Day() - time.Now().Day()
	if diff == 0 {
		return "today"
	}
	if diff == 1 {
		return "tomorrow"
	}
	if diff == -1 {
		return "yesterday"
	}
	if diff < 10 && diff > -7 {
		if diff >= 7 {
			return "next " + t.Weekday().String()
		}
		if diff > 0 {
			return t.Weekday().String()
		}
		return "last " + t.Weekday().String()
	}
	if diff > 0 {
		return "in " + fmt.Sprint(diff) + " days"
	}
	return fmt.Sprint(-diff) + " days ago"
}

var colorReset = "\033[0m"
var bold = "\033[1m"
var weak = "\033[2m"
var colorMap = map[string]string{
	"pink":    "\033[48;5;217;30m",
	"red":     "\033[48;5;203;30m",
	"orange":  "\033[48;5;208;30m",
	"yellow":  "\033[48;5;229;30m",
	"green":   "\033[48;5;120;30m",
	"blue":    "\033[48;5;45;30m",
	"purple":  "\033[48;5;147;30m",
	"brown":   "\033[48;5;101m",
	"gray":    "\033[48;5;250;30m",
	"default": "\033[48;5;240m",
}

func printTasks(tasks []Task) {
	maxClassLen := 0
	classLengths := make([]int, len(tasks))
	for i, task := range tasks {
		if len(task.Class) > maxClassLen {
			maxClassLen = len(task.Class)
		}
		classLengths[i] = len(task.Class)
	}

	for i, task := range tasks {
		classOffset := maxClassLen - classLengths[i]
		class := colorMap[task.ClassColor] + task.Class + colorReset
		name := bold + task.Name + colorReset
		due := formatRelativeDate(task.Due)
		// fmt.Printf("%s %s %s\n", class, name, due)
		fmt.Printf("%*s%s | %s  %s(%s)%s\n", classOffset, "", class, name, weak, due, colorReset)
	}
}

func main() {
	printTasks(queryNotionTaskDB())
}
