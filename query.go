package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mfigurski80/NTPeek/types"
)

type NotionEntry = types.NotionEntry

func queryNotionEntryDB(dbId string) []NotionEntry {
	// expects global variable `FieldNamesConfig: FieldNames`
	// built from command line flags (or default) for filter/sorting
	url := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", dbId)
	getBefore := time.Now().AddDate(0, 0, 9).Format("2006-01-02")
	payload := strings.NewReader(`{
		"page_size": 100,
		"filter": {
			"and": [{
				"property": "` + FieldNamesConfig.DoneField + `",
				"checkbox": {
					"equals": false
				}
			}, {
				"property": "` + FieldNamesConfig.DateField + `",
				"date": {
					"before": "` + getBefore + `"
				}
			}]
		},
		"sorts": [{
			"property": "` + FieldNamesConfig.DateField + `",
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
		panic(fmt.Errorf("Returned Error [%v]: %s\n", result["status"], result["message"]))
	}

	var entries []NotionEntry = make([]NotionEntry, len(result["results"].([]interface{})))
	for i, entry := range result["results"].([]interface{}) {
		entries[i] = entry.(map[string]interface{})["properties"].(map[string]interface{})
		entries[i]["_id"] = entry.(map[string]interface{})["id"].(string)
	}

	return entries
}

func parseNotionRichText(richText []interface{}) string {
	var text string
	for _, entry := range richText {
		text += entry.(map[string]interface{})["plain_text"].(string)
	}
	return text
}

func requireField(entry map[string]interface{}, fieldName string, fieldVal string) interface{} {
	if entry[fieldVal] == nil {
		fields := []string{}
		for k := range entry {
			fields = append(fields, k)
		}
		panic(fmt.Errorf(
			"%s field '%s' not found in query response. Found fields: %v",
			fieldName, fieldVal, fields,
		))
	}
	return entry[fieldVal]
}

type FieldNames struct {
	TitleField    string
	DoneField     string
	DateField     string
	CategoryField string
	TagField      string
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

func queryNotionTaskDB(dbId string) []Task {
	entries := queryNotionEntryDB(dbId)

	var tasks []Task
	for _, properties := range entries {
		id := properties["_id"].(string)

		requireField(properties, "title", FieldNamesConfig.TitleField)
		name := parseNotionRichText(properties[FieldNamesConfig.TitleField].(map[string]interface{})["title"].([]interface{}))

		requireField(properties, "date", FieldNamesConfig.DateField)
		due_txt := properties[FieldNamesConfig.DateField].(map[string]interface{})["date"].(map[string]interface{})["start"].(string)
		due, _ := time.ParseInLocation("2006-01-02", due_txt[:10], loc)

		requireField(properties, "category", FieldNamesConfig.CategoryField)
		class := ""
		classColor := "blue"
		cf := properties[FieldNamesConfig.CategoryField].(map[string]interface{})["select"].(map[string]interface{})
		class = cf["name"].(string)
		classColor = cf["color"].(string)

		requireField(properties, "tags", FieldNamesConfig.TagField)
		tags := []string{}
		if properties[FieldNamesConfig.TagField] != nil {
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
