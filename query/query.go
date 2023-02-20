package query

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func doNotionDBRequest(access QueryAccessArgument, sort string) (*http.Response, error) {
	// expects global variable `FieldNamesConfig: FieldNames`
	// built from command line flags (or default) for filter/sorting
	url := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", access.DBId)
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
		"sorts": ` + sort + `
	}`)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", access.Secret))
	return http.DefaultClient.Do(req)
}

func formatSortDirective(sortString SortString) string {
	// from "Due:desc,Name"	=> '[{"property": "Due", "direction": "descending"}, {"property": "Name", "direction": "ascending"}]'
	var sortDirective string = "["
	for i, prop := range strings.Split(sortString, ",") {
		sp := strings.Split(prop, ":")
		field, dir := sp[0], "ascending"
		if len(sp) > 1 {
			switch sp[1] {
			case "desc":
				dir = "descending"
			case "asc":
				dir = "ascending"
			default:
				panic(fmt.Errorf("Invalid sort direction: %s. Use 'desc', 'asc'", dir))
			}
		} else if len(sp) > 2 {
			panic(fmt.Errorf("Invalid sort string: %s. Use 'field:dir'", prop))
		}
		if i != 0 {
			sortDirective += ", "
		}
		sortDirective += fmt.Sprintf(`{"property": "%s", "direction": "%s"}`, field, dir)
	}
	sortDirective += "]"
	return sortDirective
}
