package query

import (
	"fmt"
	"net/http"
	"strings"
)

func doNotionDBRequest(access QueryAccessArgument, sort string, filter string) (*http.Response, error) {
	// sort and filter should be FORMATTED JSON
	url := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", access.DBId)

	payload := strings.NewReader(`{
		"page_size": 100,
		"filter": ` + filter + `,
		"sorts": ` + sort + `
	}`)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", access.Secret))
	return http.DefaultClient.Do(req)
}
