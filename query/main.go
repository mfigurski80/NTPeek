package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mfigurski80/NTPeek/types"
)

func QueryNotionEntryDB(secret string, dbId string) []types.NotionEntry {
	res, err := doNotionDBRequest(secret, dbId)
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
