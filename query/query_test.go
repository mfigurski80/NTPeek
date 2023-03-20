package query

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"net/http"
)

/// HTTP CLIENT MOCK

type mockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c *mockClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}

/// TESTS

var _ = Describe("function `doNotionDBRequest`", func() {
	queryAccessArg := QueryAccessArgument{"A", "B"}

	lastRequest := &http.Request{}
	httpClient = &mockClient{ // mock global client
		DoFunc: func(req *http.Request) (*http.Response, error) {
			lastRequest = req
			return &http.Response{
				StatusCode: 200,
				Body:       nil,
			}, nil
		},
	}

	BeforeEach(func() {
		lastRequest = &http.Request{}
	})

	It("should send a request with the correct query", func() {
		res, err := doNotionDBRequest(queryAccessArg, "", 100, "")
		Expect(err).To(BeNil())
		Expect(res.StatusCode).To(Equal(200))
		Expect(lastRequest.URL.String()).To(Equal("https://api.notion.com/v1/databases/B/query"))
		Expect(lastRequest.Header.Get("Authorization")).To(Equal("Bearer A"))
		Expect(lastRequest.Header.Get("Content-Type")).To(Equal("application/json"))
		Expect(lastRequest.Method).To(Equal("POST"))
	})

	It("should embed parameters correctly", func() {
		sort := `[{"property": "Due", "direction": "ascending"}]`
		filter := `{"property": "Status", "select": {"equals": "Done"}}`
		doNotionDBRequest(queryAccessArg, sort, 100, filter)
		// read lastRequest body
		b, err := ioutil.ReadAll(lastRequest.Body)
		Expect(err).To(BeNil())
		Expect(b).To(ContainSubstring(`"sorts": ` + sort))
		Expect(b).To(ContainSubstring(`"filter": ` + filter))
	})

})
