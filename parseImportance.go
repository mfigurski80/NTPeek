package main

// task importance enum type
type Importance uint8

const (
	LO Importance = iota
	AVG
	HI
)

var importanceTags = map[string]Importance{
	"exam":         HI,
	"projecttask":  HI,
	"presentation": HI,
	"project":      HI,
	"paper":        HI,
	"meeting":      LO,
	"read":         LO,
	"utility":      LO,
}

func parseImportance(task Task) Importance {
	minImportance := LO
	for _, tag := range task.Tags {
		if importance, ok := importanceTags[tag]; ok {
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
