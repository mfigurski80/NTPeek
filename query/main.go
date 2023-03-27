package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mfigurski80/NTPeek/filter"
	"github.com/mfigurski80/NTPeek/types"
)

type QueryAccessArgument struct {
	Secret string
	DBId   string
}

type QueryParamArgument struct {
	Sort   SortString
	Limit  LimitNumber
	Filter []filter.FilterString
}

func QueryNotionEntryDB(
	access QueryAccessArgument,
	param QueryParamArgument,
) ([]types.NotionEntry, error) {
	// do request
	sortDirective, err := formatSortDirective(param.Sort)
	if err != nil {
		return nil, err
	}
	limit := param.Limit
	filterDirective, err := filter.ParseFilter(param.Filter)
	if err != nil {
		return nil, err
	}
	res, err := doNotionDBRequest(access, sortDirective, limit, filterDirective)
	if err != nil {
		return nil, err
	}
	// parse into entries
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	if result["object"] == "error" {
		return nil, fmt.Errorf(
			"Returned Error [%v]: %s\n",
			result["status"], result["message"],
		)
	}

	var entries []types.NotionEntry = make(
		[]types.NotionEntry,
		len(result["results"].([]interface{})),
	)
	for i, entry := range result["results"].([]interface{}) {
		entries[i] = entry.(map[string]interface{})["properties"].(map[string]interface{})
		entries[i]["_id"] = entry.(map[string]interface{})["id"].(string)
	}
	return entries, nil
}
