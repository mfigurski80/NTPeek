package priority

import (
	"flag"
	"strings"
)

/// Tag Priority Configuration: sets up how to parse task tags

var TagsPriority = TagsPriorityMap{
	"exam":         HI,
	"projecttask":  HI,
	"presentation": HI,
	"project":      HI,
	"paper":        HI,
	"meeting":      LO,
	"read":         LO,
	"utility":      LO,
}

var DefaultPriority = MED

func SetupGlobalTagPriorityFlags(flagsets []*flag.FlagSet) func() {
	priorityTags := ""
	loPriorityTags := ""
	dPriority := true

	for _, fs := range flagsets {
		fs.StringVar(&priorityTags, "priority", "", "Comma-separated tags to render as prioritized (Default \"exam,projecttask,presentation,project,paper)\"")
		fs.StringVar(&loPriorityTags, "lo-priority", "", "Comma-separated tags to render as unprioritized (Default \"meeting,read,utility\")")
		fs.BoolVar(&dPriority, "default-priority", false, "Default avg priority if no tags are present")
	}

	return func() {
		if priorityTags != "" {
			// remove all imporant tags
			for k := range TagsPriority {
				if TagsPriority[k] == HI {
					delete(TagsPriority, k)
				}
			}
			for _, tag := range strings.Split(priorityTags, ",") {
				TagsPriority[tag] = HI
			}
		}
		if loPriorityTags != "" {
			// remove all unimportant tags
			for k := range TagsPriority {
				if TagsPriority[k] == LO {
					delete(TagsPriority, k)
				}
			}
			for _, tag := range strings.Split(loPriorityTags, ",") {
				TagsPriority[tag] = LO
			}
		}
		if dPriority {
			DefaultPriority = MED
		}
	}
}
