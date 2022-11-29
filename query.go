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
	Id         string
	Due        time.Time
	Tags       []string
	Class      string
	ClassColor string
}

var loc, _ = time.LoadLocation("Local")

func queryNotionTaskDB() []Task {
	url := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", NotionDatabaseId)
	getBefore := time.Now().AddDate(0, 0, 9).Format("2006-01-02")
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
					"before": "` + getBefore + `"
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
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", NotionAuthorizationSecret))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
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
		id := entry.(map[string]interface{})["id"].(string)

		due_txt := properties["Due"].(map[string]interface{})["date"].(map[string]interface{})["start"].(string)
		due, _ := time.ParseInLocation("2006-01-02", due_txt[:10], loc)

		class := ""
		classColor := "blue"
		if properties["Class"] != nil {
			class = properties["Class"].(map[string]interface{})["select"].(map[string]interface{})["name"].(string)
			classColor = properties["Class"].(map[string]interface{})["select"].(map[string]interface{})["color"].(string)
		}

		tags := []string{}
		if properties["Tags"] != nil {
			for _, tag := range properties["Tags"].(map[string]interface{})["multi_select"].([]interface{}) {
				// append lowercase tag
				tags = append(tags, strings.ToLower(tag.(map[string]interface{})["name"].(string)))
			}
		}

		// fmt.Printf("(%s) %s [%v] due %s\n", class, name, tags, due)
		tasks = append(tasks, Task{name, id, due, tags, class, classColor})
	}
	return tasks
}
