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
	// expects global variable `ImportanceTags: ImportanceTagsMap`
	// build from command line flags (or default)
	minImportance := LO
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
