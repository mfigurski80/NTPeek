package query

import "net/http"

type httpClientType interface {
	Do(req *http.Request) (*http.Response, error)
}

var httpClient httpClientType = http.DefaultClient
