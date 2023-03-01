package priority

import (
	"strings"

	"github.com/mfigurski80/NTPeek/types"
)

func Assign(rows []types.NotionEntry, config PriorityConfig) []Priority {
	tags := make([][]string, len(rows))
	for i, row := range rows {
		tags[i] = make([]string, 0)
		f, ok := row[config.Field].(map[string]interface{})
		if !ok {
			priorityError(config.Field, "not found in the database.")
		}
		tagResults, ok := f["multi_select"].([]interface{})
		if !ok {
			priorityError(config.Field, "is not a multi-select field")
		}
		for _, t := range tagResults {
			tags[i] = append(tags[i], t.(map[string]interface{})["name"].(string))
		}
	}
	priorities := make([]Priority, len(rows))
	for i := range rows {
		priorities[i] = parsePriority(tags[i], config)
	}
	return priorities
}

func priorityError(field string, issue string) {
	panic("Priority Assignment error: field (" + field + ") " + issue)
}

func parsePriority(tags []string, config PriorityConfig) Priority {
	minPriority := LO
	anyFound := false
	for _, tag := range tags {
		if priority, ok := config.Map[strings.ToLower(tag)]; ok {
			anyFound = true
			if priority > minPriority {
				minPriority = priority
			}
		}
	}
	if !anyFound { // if no set tags, default
		minPriority = config.Default
	}
	return minPriority
}

// TODO: OUTDATED?
func formatImportance(importance Priority, formats [3]string) string {
	return formats[importance]
}
