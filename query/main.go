package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mfigurski80/NTPeek/types"
)

type QueryAccessArgument struct {
	Secret string
	DBId   string
}

type QueryParamArgument struct {
	Sort   SortString
	Filter FilterString
}

func QueryNotionEntryDB(access QueryAccessArgument, param QueryParamArgument) []types.NotionEntry {
	// do request
	sortDirective := formatSortDirective(param.Sort)
	filterDirective := formatFilterDirective(param.Filter)
	res, err := doNotionDBRequest(access, sortDirective, filterDirective)
	if err != nil {
		panic(err)
	}
	// parse into entries
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	if result["object"] == "error" {
		panic(fmt.Errorf("Returned Error [%v]: %s\n", result["status"], result["message"]))
	}

	var entries []types.NotionEntry = make([]types.NotionEntry, len(result["results"].([]interface{})))
	for i, entry := range result["results"].([]interface{}) {
		entries[i] = entry.(map[string]interface{})["properties"].(map[string]interface{})
		entries[i]["_id"] = entry.(map[string]interface{})["id"].(string)
	}
	return entries
}
