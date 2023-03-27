package query_test

import (
	"io/ioutil"
	"net/http"

	q "github.com/mfigurski80/NTPeek/query"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sort", func() {
	queryAccessArg := q.QueryAccessArgument{"A", "B"}
	var lastRequest *http.Request
	BeforeEach(func() {
		setupHttpMock(&lastRequest)
	})

	It("correctly formats the sort directive", func() {
		sort := "Due:asc,Priority:desc"
		q.QueryNotionEntryDB(
			queryAccessArg, q.QueryParamArgument{sort, 100, []string{}})
		body, err := ioutil.ReadAll(lastRequest.Body)
		Expect(err).To(BeNil())
		Expect(string(body)).To(ContainSubstring(`"sorts": [{"property": "Due", "direction": "ascending"}, {"property": "Priority", "direction": "descending"}]`))
	})

	Context("errors in sort directive", func() {

		It("recognizes invalid sort direction", func() {
			sort := "Due:INVALID_DIRECTION"
			_, err := q.QueryNotionEntryDB(
				queryAccessArg, q.QueryParamArgument{sort, 100, []string{}})
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("direction"))
		})

		It("recognizes invalid sort syntax", func() {
			sort := "A:asc:B:C"
			_, err := q.QueryNotionEntryDB(
				queryAccessArg, q.QueryParamArgument{sort, 100, []string{}})
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("syntax"))
		})

	})

})
