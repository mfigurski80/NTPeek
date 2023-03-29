package query_test

import (
	"flag"

	q "github.com/mfigurski80/NTPeek/query"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Arguments", func() {
	It("sets up sort flag", func() {
		flagset := flag.NewFlagSet("test", 0)
		res := q.SetupSortFlag([]*flag.FlagSet{flagset})
		flagset.Parse([]string{"-sort", "TEST VALUE"})
		Expect(*res).To(Equal("TEST VALUE"))
	})
	It("sets up limit flag", func() {
		flagset := flag.NewFlagSet("test", 0)
		res := q.SetupLimitFlag([]*flag.FlagSet{flagset})
		flagset.Parse([]string{"-limit", "123"})
		Expect(*res).To(Equal(uint(123)))
	})
})
