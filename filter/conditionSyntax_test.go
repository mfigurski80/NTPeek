package filter_test

import (
	"encoding/json"

	f "github.com/mfigurski80/NTPeek/filter"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func unmarshal(s string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(s), &data)
	return data, err
}

var _ = Describe("Composite Condition Syntax", func() {

	It("unwraps simple conditions in parentheses", func() {
		r, err := f.ParseFilter([]string{"NAME:number = 1"})
		Expect(err).To(BeNil())
		r2, err := f.ParseFilter([]string{"(NAME:number = 1)"})
		Expect(err).To(BeNil())
		Expect(r).To(Equal(r2))
		r3, err := f.ParseFilter([]string{"((NAME:number = 1))"})
		Expect(err).To(BeNil())
		Expect(r).To(Equal(r3))
	})

	It("allows combining conditions with AND", func() {
		r, err := f.ParseFilter([]string{`NAME:number = 1 AND NAME:number = 1`})
		Expect(err).To(BeNil())
		data, err := unmarshal(r)
		Expect(err).To(BeNil())
		Expect(data).To(HaveKey("and"))
		By("allowing multiple ANDs")
		r2, err := f.ParseFilter([]string{`NAME:number = 1 AND NAME:number = 1 AND NAME:number = 1`})
		Expect(err).To(BeNil())
		data2, err := unmarshal(r2)
		Expect(err).To(BeNil())
		Expect(data2).To(HaveKey("and"))
		Expect(data2["and"]).To(HaveLen(3))
	})

	It("allows combining conditions with OR", func() {
		r, err := f.ParseFilter([]string{`NAME:number = 1 OR NAME:number = 1`})
		Expect(err).To(BeNil())
		data, err := unmarshal(r)
		Expect(err).To(BeNil())
		Expect(data).To(HaveKey("or"))
		By("allowing multiple ORs")
		r2, err := f.ParseFilter([]string{`NAME:number = 1 OR NAME:number = 1 OR NAME:number = 1`})
		Expect(err).To(BeNil())
		data2, err := unmarshal(r2)
		Expect(err).To(BeNil())
		Expect(data2).To(HaveKey("or"))
		Expect(data2["or"]).To(HaveLen(3))
	})

	It("allows combining conditions with AND and OR", func() {
		r, err := f.ParseFilter([]string{`NAME:number = 1 AND (NAME:number = 1 OR NAME:number = 1)`})
		Expect(err).To(BeNil())
		data, err := unmarshal(r)
		Expect(err).To(BeNil())
		Expect(data).To(HaveKey("and"))
		Expect(data["and"]).To(HaveLen(2))
		Expect(data["and"].([]interface{})[1]).To(HaveKey("or"))
	})

})
