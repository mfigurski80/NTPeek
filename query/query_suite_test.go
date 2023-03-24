package query_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	q "github.com/mfigurski80/NTPeek/query"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestQuery(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Query Suite")
}

/// GLOBAL HTTP-CLIENT MOCK

type mockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c *mockClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}

func setupHttpMock(last **http.Request) {
	q.SET_HTTP_CLIENT(&mockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			*last = req
			return &http.Response{
				StatusCode: 200,
				Body: ioutil.NopCloser(bytes.NewBufferString(
					`{"object": "list", "results": []}`,
				)),
			}, nil
		},
	})
}

/// MAIN SPEC TESTS

var _ = Describe("`QueryNotionEntryDB` function", func() {
	queryAccessArg := q.QueryAccessArgument{"A", "B"}
	var lastRequest *http.Request
	BeforeEach(func() {
		lastRequest = nil
	})

	Context("while the query is successful", func() {

		BeforeEach(func() {
			setupHttpMock(&lastRequest)
		})

		It("should not error", func() {
			res, err := q.QueryNotionEntryDB(
				queryAccessArg, q.QueryParamArgument{"", 100, []string{}})
			Expect(err).To(BeNil())
			Expect(res).ToNot(BeNil())
		})

		It("should make correct request", func() {
			q.QueryNotionEntryDB(
				queryAccessArg, q.QueryParamArgument{"", 100, []string{}})
			Expect(lastRequest.URL.String()).To(Equal(
				"https://api.notion.com/v1/databases/B/query"))
			Expect(lastRequest.Header.Get("Authorization")).To(Equal("Bearer A"))
			Expect(lastRequest.Header.Get("Content-Type")).To(Equal("application/json"))
			Expect(lastRequest.Method).To(Equal("POST"))
		})

		It("correctly formats the filter directive", func() {
			filter := "(Due:date > 2021/01/01 AND (Done:checkbox = FALSE OR Due:date = EMPTY))"
			_, err := q.QueryNotionEntryDB(
				queryAccessArg, q.QueryParamArgument{"", 100, []string{filter}})
			Expect(err).To(BeNil())
			body, err := ioutil.ReadAll(lastRequest.Body)
			Expect(err).To(BeNil())
			Expect(string(body)).To(ContainSubstring(`"filter": {"and": [{"property": "Due", "date": {"after": "2021-01-01"}}, {"or": [{"property": "Done", "checkbox": {"equals": false}}, {"property": "Due", "date": {"is_empty": true}}]}]}`))
		})

		It("correctly limits with the page size field", func() {
			_, err := q.QueryNotionEntryDB(
				queryAccessArg, q.QueryParamArgument{"", 10, []string{}})
			Expect(err).To(BeNil())
			body, err := ioutil.ReadAll(lastRequest.Body)
			Expect(err).To(BeNil())
			Expect(string(body)).To(ContainSubstring(`"page_size": 10`))
		})

	})

	Context("while the query fails", func() {

		BeforeEach(func() {
			q.SET_HTTP_CLIENT(&mockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					lastRequest = req
					return &http.Response{
						StatusCode: 400,
						Body: ioutil.NopCloser(bytes.NewBufferString(
							`{"object": "error", "status": 400, "code": "invalid_request", "message": "ERROR MESSAGE"}`,
						)),
					}, nil
				},
			})
		})

		It("should surface the error", func() {
			_, err := q.QueryNotionEntryDB(
				queryAccessArg, q.QueryParamArgument{"", 100, []string{}})
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("ERROR MESSAGE"))
			Expect(err.Error()).To(ContainSubstring("400"))
		})

	})

})
