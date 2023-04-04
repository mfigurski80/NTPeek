package filter_test

import (
	"testing"
	"time"

	f "github.com/mfigurski80/NTPeek/filter"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

func TestFilter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filter Suite")
}

/// MAIN SPEC TESTS

var _ = Describe("Filter", func() {

	testFilters := []string{
		`NAME:text = "VAL" OR DATE:date < NEXT 10 DAY`,
		`DONE:checkbox = true AND NUM:number > 10`,
	}

	It("runs benchmarks", Serial, Label("benchmark"), func() {
		experiment := gmeasure.NewExperiment("format filters")
		AddReportEntry(experiment.Name, experiment)
		experiment.Sample(func(idx int) {
			experiment.MeasureDuration("ParseFilter", func() {
				for i := 0; i < 10; i++ {
					f.ParseFilter(testFilters)
				}
			})
		}, gmeasure.SamplingConfig{N: 2000, Duration: time.Minute})
		stats := experiment.GetStats("format filters")
		mean := stats.DurationFor(gmeasure.StatMean)
		median := stats.DurationFor(gmeasure.StatMedian)
		Expect(mean).To(BeNumerically("~", median, 0.1), "mean and median should be close")
	})

})
