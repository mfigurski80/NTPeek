package render_test

import (
	"encoding/json"
	"testing"

	"github.com/mfigurski80/NTPeek/priority"
	r "github.com/mfigurski80/NTPeek/render"
	"github.com/mfigurski80/NTPeek/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRender(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Render Suite")
}

/// MAIN SPEC TESTS

var _ = Describe("`RenderTasks` function", func() {
	var defaultTask types.NotionEntry
	var defaultPriorityConfig = priority.PriorityConfig{
		Field:   "MULTI_SELECT FIELD",
		Map:     priority.TagsPriorityMap{},
		Default: priority.MED,
	}

	BeforeEach(func() {
		err := json.Unmarshal([]byte(ntResponse), &defaultTask)
		Expect(err).To(BeNil())
	})

	Context("with a valid template", func() {
		It("should accept a valid template", func() {
			r.RenderTasks([]types.NotionEntry{defaultTask}, "...", defaultPriorityConfig)
		})
	})

})

// POST Nt Response `properties`: to be parsed into a NotionEntry
// From https://developers.notion.com/reference/post-database-query
const ntResponse = `{
"SELECT FIELD": {
	"type": "select",
	"select": {
		"name": "SELECT VALUE",
		"color": "pink"
	}
},
"MULTI_SELECT FIELD": {
	"type": "multi_select",
	"multi_select": [{
		"name": "MULTI_SELECT VALUE 1",
		"color": "pink"
	}, {
		"name": "MULTI_SELECT VALUE 2",
		"color": "pink"
	}]
},
"NUMBER FIELD": {
	"type": "number",
	"number": 123
},
"DATE FIELD": {
	"type": "date",
	"date": { "start": "2021-01-01" }
},
"TEXT FIELD": {
	"type": "rich_text",
	"rich_text": [{
		"type": "text",
		"text": { "content": "TEXT VALUE" }
	}]
}}`
