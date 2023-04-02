package render_test

import (
	"encoding/json"

	"github.com/acarl005/stripansi"
	"github.com/mfigurski80/NTPeek/priority"
	r "github.com/mfigurski80/NTPeek/render"
	"github.com/mfigurski80/NTPeek/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Render Generic Values", func() {
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

	It("should render select fields", func() {
		s, err := r.RenderTasks(defaultTasks, "%SELECT FIELD%", defaultPriorityConfig)
		Expect(err).To(BeNil())
		Expect(stripansi.Strip(s)).To(Equal("SELECT VALUE\n"))
	})

	It("should render title fields", func() {
		s, err := r.RenderTasks(defaultTasks, "%TITLE FIELD%", defaultPriorityConfig)
		Expect(err).To(BeNil())
		Expect(s).To(Equal("TITLE VALUE\n"))
	})

	It("should render text fields", func() {
		s, err := r.RenderTasks(defaultTasks, "%TEXT FIELD%", defaultPriorityConfig)
		Expect(err).To(BeNil())
		Expect(s).To(ContainSubstring("TEXT VALUE"))
	})

	It("should render multi-select fields", func() {
		s, err := r.RenderTasks(defaultTasks, "%MULTI_SELECT FIELD%", defaultPriorityConfig)
		Expect(err).To(BeNil())
		Expect(s).To(ContainSubstring("MULTI_SELECT VALUE 1"))
		Expect(s).To(ContainSubstring("MULTI_SELECT VALUE 2"))
	})

	It("should render number fields", func() {
		s, err := r.RenderTasks(defaultTasks, "%NUMBER FIELD%", defaultPriorityConfig)
		Expect(err).To(BeNil())
		Expect(stripansi.Strip(s)).To(Equal("123\n"))
	})

	When("Rendering date fields", func() {
		It("should render relative dateS by default", func() {
			s, err := r.RenderTasks(defaultTasks, "%DATE FIELD%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(s).To(ContainSubstring("days ago"))
		})

		It("should render absolute dates", func() {
			s, err := r.RenderTasks(defaultTasks, "%DATE FIELD:full%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(s).To(ContainSubstring("2021-01-01"))
		})

		It("should render simple dates", func() {
			s, err := r.RenderTasks(defaultTasks, "%DATE FIELD:simple%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(s).To(ContainSubstring("Jan 1"))
		})

		It("should render relative dates", func() {
			s, err := r.RenderTasks(defaultTasks, "%DATE FIELD:relative%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(s).To(ContainSubstring("days ago"))
		})
	})

	It("should render checkbox fields", func() {
		s, err := r.RenderTasks(defaultTasks, "%CHECKBOX FIELD%", defaultPriorityConfig)
		Expect(err).To(BeNil())
		GinkgoT().Log(s)
		Expect(s).To(ContainSubstring("[x]"))
	})

	When("rendering ID _id fields", func() {
		It("should render full by default", func() {
			s, err := r.RenderTasks(defaultTasks, "%_id%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(s).To(ContainSubstring("1234-5678-9012-3456"))
		})
		It("should render short with modifier", func() {
			s, err := r.RenderTasks(defaultTasks, "%_id:short%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(s).To(ContainSubstring("1234"))
		})
	})

	When("rendering priority _p fields", func() {
		It("should render empty by default", func() {
			s, err := r.RenderTasks(defaultTasks, "%_p%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(s).To(ContainSubstring(" "))
		})
		It("should render value if priority", func() {
			conf := priority.PriorityConfig{
				Field:   "MULTI_SELECT FIELD",
				Map:     priority.TagsPriorityMap{},
				Default: priority.HI,
			}
			s, err := r.RenderTasks(defaultTasks, "%_p%", conf)
			Expect(err).To(BeNil())
			Expect(s).To(ContainSubstring("!"))
		})
	})

	Context("rendering a mis-formatted field", func() {
		It("should throw an internal error", func() {
			rNames := []string{"title", "rich_text", "select", "multi_select", "number", "date", "checkbox"}
			for _, rName := range rNames {
				task := types.NotionEntry{
					"MULTI_SELECT FIELD": map[string]interface{}{"type": "multi_select", "multi_select": []interface{}{}},
					rName:                map[string]interface{}{"type": rName, "foo": "bar"},
				}
				_, err := r.RenderTasks([]types.NotionEntry{task}, "%"+rName+"%", defaultPriorityConfig)
				Expect(err).ToNot(BeNil(), "Expected error for "+rName)
				Expect(err.Error()).To(ContainSubstring("internal"), "Expected internal error for "+rName)
			}
		})
	})

})
