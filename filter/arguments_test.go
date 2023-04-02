package filter_test

import (
	"flag"

	f "github.com/mfigurski80/NTPeek/filter"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Arguments", func() {
	It("sets up filter flag", func() {
		flagset := flag.NewFlagSet("test", flag.ContinueOnError)
		filterFlags := f.SetupFilterFlag([]*flag.FlagSet{flagset})
		flagset.Parse([]string{"-filter", "TEST VALUE"})
		Expect(*filterFlags).To(HaveLen(1))
		Expect((*filterFlags)[0]).To(Equal("TEST VALUE"))
	})
})
