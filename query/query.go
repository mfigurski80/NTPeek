package query

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/mfigurski80/NTPeek/types"
)

type NotionEntry = types.NotionEntry

func doNotionDBRequest(secret string, dbId string) (*http.Response, error) {
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
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", secret))
	return http.DefaultClient.Do(req)
}
