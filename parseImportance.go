package main

// task importance enum type
type Importance uint8

const (
	LO Importance = iota
	AVG
	HI
)

type ImportanceTagsMap map[string]Importance

func parseImportance(task Task) Importance {
	// expects global variables `ImportanceTags: ImportanceTagsMap`,
	// and `DefaultTagImportance: Importance` build from command
	// line flags (or default)
	minImportance := LO
	if len(task.Tags) == 0 {
		return DefaultImportance
	}
	for _, tag := range task.Tags {
		if importance, ok := ImportanceTags[tag]; ok {
			if importance > minImportance {
				minImportance = importance
			}
		} else {
			minImportance = AVG
		}
	}
	return minImportance
}

func formatImportance(importance Importance, formats [3]string) string {
	return formats[importance]
}
