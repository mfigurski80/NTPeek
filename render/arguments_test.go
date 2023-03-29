package render_test

import (
	"flag"

	r "github.com/mfigurski80/NTPeek/render"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Arguments", func() {
	It("sets up select flag", func() {
		flagset := flag.NewFlagSet("test", flag.ExitOnError)
		res := r.SetupSelectFlag([]*flag.FlagSet{flagset})
		flagset.Parse([]string{"-select", "TEST VALUE"})
		Expect(*res).To(Equal("TEST VALUE"))
	})
})
