package filter_test

import (
	"fmt"
	"time"

	f "github.com/mfigurski80/NTPeek/filter"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Simple Filter Syntax", func() {

	Context("describing text field", func() {
		It("should support equality", func() {
			r, err := f.ParseFilter([]string{`NAME:text = "VAL"`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "title": {"equals": "VAL"}}`))
		})
		It("should support inequality", func() {
			r, err := f.ParseFilter([]string{`NAME:text != "VAL"`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "title": {"does_not_equal": "VAL"}}`))
		})
		It("should support contains", func() {
			r, err := f.ParseFilter([]string{`NAME:text CONTAINS "VAL"`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "title": {"contains": "VAL"}}`))
		})
		It("should support not contains", func() {
			r, err := f.ParseFilter([]string{`NAME:text !CONTAINS "VAL"`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "title": {"does_not_contain": "VAL"}}`))
		})
		It("should support starts with", func() {
			r, err := f.ParseFilter([]string{`NAME:text STARTS_WITH "VAL"`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "title": {"starts_with": "VAL"}}`))
		})
		It("should support ends with", func() {
			r, err := f.ParseFilter([]string{`NAME:text ENDS_WITH "VAL"`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "title": {"ends_with": "VAL"}}`))
		})
		It("should list supported operators", func() {
			supported := []string{"=", "!=", "CONTAINS", "!CONTAINS", "STARTS_WITH", "ENDS_WITH"}
			_, err := f.ParseFilter([]string{`NAME:text > "VAL"`})
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("support"))
			for _, s := range supported {
				Expect(err.Error()).To(ContainSubstring(s))
			}
		})
	})

	Context("describing select field", func() {
		It("should support equality", func() {
			r, err := f.ParseFilter([]string{`NAME:select = "VAL"`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "select": {"equals": "VAL"}}`))
		})
		It("should support inequality", func() {
			r, err := f.ParseFilter([]string{`NAME:select != "VAL"`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "select": {"does_not_equal": "VAL"}}`))
		})
		It("should list supported operators", func() {
			supported := []string{"=", "!="}
			_, err := f.ParseFilter([]string{`NAME:select > "VAL"`})
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("support"))
			for _, s := range supported {
				Expect(err.Error()).To(ContainSubstring(s))
			}
		})
	})

	Context("describing checkbox field", func() {
		It("should support bool equality", func() {
			r, err := f.ParseFilter([]string{`NAME:checkbox = TRUE`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "checkbox": {"equals": true}}`))
			r, err = f.ParseFilter([]string{`NAME:checkbox = FALSE`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "checkbox": {"equals": false}}`))
		})
		It("should support bool inequality", func() {
			r, err := f.ParseFilter([]string{`NAME:checkbox != TRUE`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "checkbox": {"does_not_equal": true}}`))
			r, err = f.ParseFilter([]string{`NAME:checkbox != FALSE`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "checkbox": {"does_not_equal": false}}`))
		})
		It("should list supported operators", func() {
			supported := []string{"=", "!="}
			_, err := f.ParseFilter([]string{`NAME:checkbox > "VAL"`})
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("support"))
			for _, s := range supported {
				Expect(err.Error()).To(ContainSubstring(s))
			}
		})
	})

	Context("describing number field", func() {
		It("should support equality", func() {
			r, err := f.ParseFilter([]string{`NAME:number = 1`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "number": {"equals": 1.000000}}`))
		})
		It("should support inequality", func() {
			r, err := f.ParseFilter([]string{`NAME:number != 1`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "number": {"does_not_equal": 1.000000}}`))
		})
		It("should support less than", func() {
			r, err := f.ParseFilter([]string{`NAME:number < 1`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "number": {"less_than": 1.000000}}`))
		})
		It("should support greater than", func() {
			r, err := f.ParseFilter([]string{`NAME:number > 1`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "number": {"greater_than": 1.000000}}`))
		})
		It("should support less than or equal", func() {
			r, err := f.ParseFilter([]string{`NAME:number <= 1`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "number": {"less_than_or_equal_to": 1.000000}}`))
		})
		It("should support greater than or equal", func() {
			r, err := f.ParseFilter([]string{`NAME:number >= 1`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "number": {"greater_than_or_equal_to": 1.000000}}`))
		})
		It("should list supported operators", func() {
			supported := []string{"=", "!=", "<", ">", "<=", ">="}
			_, err := f.ParseFilter([]string{`NAME:number CONTAINS "VAL"`})
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("support"))
			for _, s := range supported {
				Expect(err.Error()).To(ContainSubstring(s))
			}
		})
	})

	Context("describing multiselect field", func() {
		It("should support contains", func() {
			r, err := f.ParseFilter([]string{`NAME:multiselect CONTAINS "VAL"`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "multi-select": {"contains": "VAL"}}`))
		})
		It("should support does not contain", func() {
			r, err := f.ParseFilter([]string{`NAME:multiselect !CONTAINS "VAL"`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "multi-select": {"does_not_contain": "VAL"}}`))
		})
		It("should list supported operators", func() {
			supported := []string{"CONTAINS", "!CONTAINS"}
			_, err := f.ParseFilter([]string{`NAME:multiselect = "VAL"`})
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("support"))
			for _, s := range supported {
				Expect(err.Error()).To(ContainSubstring(s))
			}
		})
	})

	Context("describing date field", func() {
		When("using a relative date", func() {
			BeforeEach(func() {
				f.SET_TIME_PROVIDER(func() time.Time {
					v, _ := time.Parse("2006-01-02", "2020-01-01")
					return v
				})
			})
			AfterEach(func() {
				f.SET_TIME_PROVIDER(time.Now)
			})
			It("should support equality", func() {
				r, err := f.ParseFilter([]string{`NAME:date = NEXT 1 DAY`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"equals": "2020-01-02"}}`))
			})
			It("should support after", func() {
				r, err := f.ParseFilter([]string{`NAME:date > NEXT 1 DAY`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"after": "2020-01-02"}}`))
			})
			It("should support on or after", func() {
				r, err := f.ParseFilter([]string{`NAME:date >= NEXT 1 DAY`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"on_or_after": "2020-01-02"}}`))
			})
			It("should support before", func() {
				r, err := f.ParseFilter([]string{`NAME:date < NEXT 1 DAY`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"before": "2020-01-02"}}`))
			})
			It("should support on or before", func() {
				r, err := f.ParseFilter([]string{`NAME:date <= NEXT 1 DAY`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"on_or_before": "2020-01-02"}}`))
			})
			It("should parse all supported units", func() {
				r, err := f.ParseFilter([]string{`NAME:date = NEXT 1 DAY`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"equals": "2020-01-02"}}`))
				r, err = f.ParseFilter([]string{`NAME:date = NEXT 1 WEEK`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"equals": "2020-01-08"}}`))
				r, err = f.ParseFilter([]string{`NAME:date = NEXT 1 MONTH`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"equals": "2020-02-01"}}`))
				r, err = f.ParseFilter([]string{`NAME:date = NEXT 1 YEAR`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"equals": "2021-01-01"}}`))
			})
			It("should support NEXT/LAST modifiers", func() {
				r, err := f.ParseFilter([]string{`NAME:date = NEXT 1 DAY`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"equals": "2020-01-02"}}`))
				r, err = f.ParseFilter([]string{`NAME:date = LAST 1 MONTH`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"equals": "2019-12-01"}}`))
			})
			It("should support a unit amount, and default to 1", func() {
				r, err := f.ParseFilter([]string{`NAME:date = NEXT DAY`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"equals": "2020-01-02"}}`))
				r, err = f.ParseFilter([]string{`NAME:date = NEXT 2 DAY`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"equals": "2020-01-03"}}`))
			})
		})
		When("using an absolute date", func() {
			It("should support equality", func() {
				r, err := f.ParseFilter([]string{`NAME:date = 2020/01/01`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"equals": "2020-01-01"}}`))
			})
			It("should support after", func() {
				r, err := f.ParseFilter([]string{`NAME:date > 2020/01/01`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"after": "2020-01-01"}}`))
			})
			It("should support on or after", func() {
				r, err := f.ParseFilter([]string{`NAME:date >= 2020/01/01`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"on_or_after": "2020-01-01"}}`))
			})
			It("should support before", func() {
				r, err := f.ParseFilter([]string{`NAME:date < 2020/01/01`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"before": "2020-01-01"}}`))
			})
			It("should support on or before", func() {
				r, err := f.ParseFilter([]string{`NAME:date <= 2020/01/01`})
				Expect(err).To(BeNil())
				Expect(r).To(Equal(`{"property": "NAME", "date": {"on_or_before": "2020-01-01"}}`))
			})
		})
		It("should list supported operators", func() {
			supported := []string{"=", "<", ">", "<=", ">="}
			_, err := f.ParseFilter([]string{`NAME:date CONTAINS "VAL"`})
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("support"))
			for _, s := range supported {
				Expect(err.Error()).To(ContainSubstring(s))
			}
		})
	})

	Context("querying for empty fields", func() {
		It("should format with unique operator", func() {
			r, err := f.ParseFilter([]string{`NAME:text = EMPTY`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "title": {"is_empty": true}}`))
		})
		It("should allow for NOT", func() {
			r, err := f.ParseFilter([]string{`NAME:text != EMPTY`})
			Expect(err).To(BeNil())
			Expect(r).To(Equal(`{"property": "NAME", "title": {"is_not_empty": true}}`))
		})
		It("Should allow supported types", func() {
			t := []string{"text", "number", "select", "multiselect", "date"}
			for _, s := range t {
				_, err := f.ParseFilter([]string{fmt.Sprintf(`NAME:%s = EMPTY`, s)})
				Expect(err).To(BeNil(), fmt.Sprintf("type %s should be supported", s))
			}
		})
		It("should disallow unsupported types", func() {
			Skip("feature is not implemented")
			t := []string{"checkbox"}
			for _, s := range t {
				_, err := f.ParseFilter([]string{fmt.Sprintf(`NAME:%s = EMPTY`, s)})
				Expect(err).ToNot(BeNil(), fmt.Sprintf("type %s should not be supported", s))
				Expect(err.Error()).To(ContainSubstring("not supported"))
			}
		})
	})

	Context("negating operators", func() {
		It("should accept supported operators", func() {
			supportedOps := []string{"=", "CONTAINS"}
			for _, s := range supportedOps {
				_, err := f.ParseFilter([]string{fmt.Sprintf(`NAME:text ! %s "VAL"`, s)})
				Expect(err).To(BeNil(), fmt.Sprintf("operator %s should be supported", s))
			}
		})
		It("should disallow unsupported operators", func() {
			unSupportedOps := []string{"<", ">", "<=", ">="}
			for _, s := range unSupportedOps {
				_, err := f.ParseFilter([]string{fmt.Sprintf(`NAME:number ! %s 1`, s)})
				Expect(err).ToNot(BeNil(), fmt.Sprintf("operator %s should not be supported", s))
				Expect(err.Error()).To(ContainSubstring("negate"))
			}
		})
		It("! and NOT are interchangeable", func() {
			r1, err := f.ParseFilter([]string{`NAME:text ! = "VAL"`})
			Expect(err).To(BeNil())
			r2, err := f.ParseFilter([]string{`NAME:text NOT = "VAL"`})
			Expect(err).To(BeNil())
			Expect(r1).To(Equal(r2))
		})
	})

	It("provides help text for bad types", func() {
		Skip("feature is not implemented")
		_, err := f.ParseFilter([]string{`NAME:badtype = "VAL"`})
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("badtype"))
		Expect(err.Error()).To(ContainSubstring("supported types"))
	})

})
