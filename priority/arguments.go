package priority

import (
	"flag"
	"strings"
)

/// Tag Priority Configuration: sets up how to parse task tags

func SetupPriorityFlags(flagsets []*flag.FlagSet) func() PriorityConfig {
	priorityField := ""
	priorityTags := ""
	loPriorityTags := ""
	dPriority := ""

	for _, fs := range flagsets {
		fs.StringVar(&priorityField, "priority-field", "Tags",
			"Multiselect field to look at for priority tags")
		fs.StringVar(&priorityTags, "priority", "exam,projecttask,presentation,project,paper",
			"Multiselect field and comma-separated tags to render as prioritized")
		fs.StringVar(&loPriorityTags, "priority-lo", "meeting,read,utility",
			"Comma-separated tags to render as unprioritized")
		fs.StringVar(&dPriority, "priority-default", "med",
			"Default priority no tags are present (low, med, high)")
	}

	return func() PriorityConfig {
		// init with defaults
		conf := PriorityConfig{
			Field:   priorityField,
			Map:     TagsPriorityMap{},
			Default: MED,
		}
		// build map
		for _, tag := range strings.Split(loPriorityTags, ",") {
			conf.Map[strings.ToLower(tag)] = LO
		}
		for _, tag := range strings.Split(priorityTags, ",") {
			conf.Map[strings.ToLower(tag)] = HI
		}
		// set default
		switch strings.ToLower(dPriority) {
		case "low", "lo":
			conf.Default = LO
		case "high", "hi":
			conf.Default = HI
		}
		return conf
	}
}
