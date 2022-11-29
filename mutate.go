package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func mutateNotionMarkTaskDone(taskId string) error {
	url := fmt.Sprintf("https://api.notion.com/v1/pages/%s", NotionDatabaseId)

	payload := strings.NewReader(`{
		"properties": {
			"Done": {
				"checkbox": true
			}
		}
	}`)
	req, _ := http.NewRequest("PATCH", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", NotionAuthorizationSecret))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(string(body))
	}

	return nil
}
