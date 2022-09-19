package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
		class := properties["Class"].(map[string]interface{})["select"].(map[string]interface{})["name"].(string)
		due := properties["Due"].(map[string]interface{})["date"].(map[string]interface{})["start"].(string)
		// fmt.Printf("(%s) %s due %s\n", class, name, due)
		tasks = append(tasks, Task{name, due, class, ""})
	}
	return tasks
}

func printTasks(tasks []Task) {
	maxClassLen := 0
	for _, task := range tasks {
		if len(task.Class) > maxClassLen {
			maxClassLen = len(task.Class)
		}
	}
	for _, task := range tasks {
		fmt.Printf("%*s | %s (due %s)\n", maxClassLen, task.Class, task.Name, task.Due)
	}
}

func main() {
	printTasks(queryNotionTaskDB())
}
