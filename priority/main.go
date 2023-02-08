package priority

type Priority uint8

const (
	LO Priority = iota
	MED
	HI
)

type TagsPriorityMap map[string]Priority

func ParsePriority(tags []string) Priority {
	// expects global variables `TagsPriority: TagsPriorityMap`,
	// and `DefaultPriority: Priority` build from command line flags
	minPriority := LO
	anyFound := false
	for _, tag := range tags {
		if priority, ok := TagsPriority[tag]; ok {
			anyFound = true
			if priority > minPriority {
				minPriority = priority
			}
		}
	}
	if !anyFound { // if no set tags, default
		minPriority = DefaultPriority
	}
	return minPriority
}

// TODO: OUTDATED?
func formatImportance(importance Priority, formats [3]string) string {
	return formats[importance]
}
