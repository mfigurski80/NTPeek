package priority

import (
	"strings"

	"github.com/mfigurski80/NTPeek/types"
)

type Priority uint8

const (
	LO Priority = iota
	MED
	HI
)

func Assign(rows []types.NotionEntry) []Priority {
	tags := make([][]string, len(rows))
	for i, row := range rows {
		tags[i] = make([]string, 0)
		// TODO: assumes 'Tags' is a multi-select
		for _, t := range row["Tags"].(map[string]interface{})["multi_select"].([]interface{}) {
			tags[i] = append(tags[i], t.(map[string]interface{})["name"].(string))
		}
	}
	priorities := make([]Priority, len(rows))
	for i := range rows {
		priorities[i] = ParsePriority(tags[i])
	}
	return priorities
}

type TagsPriorityMap map[string]Priority

func ParsePriority(tags []string) Priority {
	// expects global variables `TagsPriority: TagsPriorityMap`,
	// and `DefaultPriority: Priority` build from command line flags
	minPriority := LO
	anyFound := false
	for _, tag := range tags {
		if priority, ok := TagsPriority[strings.ToLower(tag)]; ok {
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
