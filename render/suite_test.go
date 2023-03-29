package render_test

import (
	"encoding/json"
	"testing"

	"github.com/acarl005/stripansi"
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
	var defaultTasks []types.NotionEntry
	var defaultPriorityConfig = priority.PriorityConfig{
		Field:   "MULTI_SELECT FIELD",
		Map:     priority.TagsPriorityMap{},
		Default: priority.MED,
	}

	BeforeEach(func() {
		t := types.NotionEntry{}
		err := json.Unmarshal([]byte(ntResponse), &t)
		Expect(err).To(BeNil())
		defaultTasks = []types.NotionEntry{t}
	})

	Context("with a valid template", func() {
		It("should render a simple literal template", func() {
			Expect(r.RenderTasks(defaultTasks, "xxx", defaultPriorityConfig)).
				Should(Equal("xxx\n"))
		})
		It("should accept single-field template", func() {
			s, err := r.RenderTasks(defaultTasks, "%UNDERSCORE_FIELD%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(stripansi.Strip(s)).To(Equal("UNDERSCORE_VALUE\n"))
		})
		It("should accept multi-field template", func() {
			s, err := r.RenderTasks(defaultTasks, "%UNDERSCORE_FIELD%%UNDERSCORE_FIELD%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(stripansi.Strip(s)).To(Equal("UNDERSCORE_VALUEUNDERSCORE_VALUE\n"))
		})
		It("should accept fields with format text", func() {
			s, err := r.RenderTasks(defaultTasks, "||%UNDERSCORE_FIELD%||", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(stripansi.Strip(s)).To(Equal("||UNDERSCORE_VALUE||\n"))
		})
		It("should accept fields with spaces", func() {
			s, err := r.RenderTasks(defaultTasks, "||%SELECT FIELD%||", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(stripansi.Strip(s)).To(Equal("||SELECT VALUE||\n"))
		})
		It("should accept modifiers", func() {
			s, err := r.RenderTasks(defaultTasks, "%SELECT FIELD:left%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(s).To(Not(ContainSubstring("ERR")))
		})
	})

	Context("with an invalid template", func() {
		It("should throw error on invalid fields", func() {
			_, err := r.RenderTasks(defaultTasks, "%INV%", defaultPriorityConfig)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("ERR"))
		})
		It("should list valid fields on invalid field error", func() {
			_, err := r.RenderTasks(defaultTasks, "%INV%", defaultPriorityConfig)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("SELECT FIELD"))
			Expect(err.Error()).To(ContainSubstring("MULTI_SELECT FIELD"))
			Expect(err.Error()).To(ContainSubstring("NUMBER FIELD"))
			Expect(err.Error()).To(ContainSubstring("DATE FIELD"))
			Expect(err.Error()).To(ContainSubstring("TITLE FIELD"))
			Expect(err.Error()).To(ContainSubstring("CHECKBOX FIELD"))
			Expect(err.Error()).To(ContainSubstring("UNDERSCORE_FIELD"))
		})
		It("should throw error on invalid modifiers", func() {
			Skip("TODO: fix this test")
			_, err := r.RenderTasks(defaultTasks, "%UNDERSCORE_FIELD:inv%", defaultPriorityConfig)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("ERR"))
		})
		It("should list valid modifiers on invalid modifier error", func() {
			Skip("TODO: fix this test")
			_, err := r.RenderTasks(defaultTasks, "%UNDERSCORE_FIELD:inv%", defaultPriorityConfig)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("left"))
			Expect(err.Error()).To(ContainSubstring("right"))
			Expect(err.Error()).To(ContainSubstring("center"))
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
	"TITLE FIELD": {
		"type": "title",
		"title": [{
			"type": "text",
			"text": { "content": "TITLE VALUE" },
			"plain_text": "TITLE VALUE"
		}]
	},
	"TEXT FIELD": {
		"type": "rich_text",
		"rich_text": [{
			"type": "text",
			"text": { "content": "TEXT VALUE" },
			"plain_text": "TEXT VALUE"
		}]
	},
	"CHECKBOX FIELD": {
		"type": "checkbox",
		"checkbox": true
	},
	"UNDERSCORE_FIELD": {
		"type": "select",
		"select": {
			"name": "UNDERSCORE_VALUE",
			"color": "pink"
		}
	},
	"_id": "1234-5678-9012-3456"
}`
