package render_test

import (
	"encoding/json"
	"strings"

	"github.com/acarl005/stripansi"
	"github.com/mfigurski80/NTPeek/priority"
	r "github.com/mfigurski80/NTPeek/render"
	"github.com/mfigurski80/NTPeek/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("global render modifiers", func() {

	var defaultTasks []types.NotionEntry
	err := json.Unmarshal([]byte(modifierTestResponse), &defaultTasks)
	if err != nil {
		panic(err)
	}
	var defaultPriorityConfig = priority.PriorityConfig{
		Field:   "name",
		Map:     priority.TagsPriorityMap{},
		Default: priority.MED,
	}

	Context("alignment modifiers", func() {

		It("properly handles `left` modifier", func() {
			res, err := r.RenderTasks(defaultTasks, "%name:left%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(stripansi.Strip(res)).To(ContainSubstring("VAL             \n"))
		})

		It("properly handles the `right` modifier", func() {
			res, err := r.RenderTasks(defaultTasks, "%name:right%", defaultPriorityConfig)
			Expect(err).To(BeNil())
			Expect(stripansi.Strip(res)).To(ContainSubstring("            VAL\n"))
		})

		It("properly handles the `center` modifier", func() {
			Skip("not implemented")
		})

	})

	It("properly handles the `bold` modifier", func() {
		res, err := r.RenderTasks(defaultTasks, "%name:bold%", defaultPriorityConfig)
		Expect(err).To(BeNil())
		Expect(res).ToNot(Equal(stripansi.Strip(res)))
		Expect(strings.Split(res, "\n")).To(HaveLen(3))
		Expect(strings.Split(res, "\n")[0]).To(ContainSubstring("\x1b"))
		Expect(strings.Split(res, "\n")[1]).To(ContainSubstring("\x1b"))
	})

})

const modifierTestResponse = `[
	{ "name": {
		"type": "multi_select",
		"multi_select": [{
			"name": "VAL",
			"color": "pink"
		}]
	}},
	{ "name": {
		"type": "multi_select",
		"multi_select": [{
			"name": "_THIS_LEN_IS_16_",
			"color": "pink"
		}]
	}}
]`
