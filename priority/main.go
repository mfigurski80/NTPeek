package priority

import (
	"fmt"
	"strings"

	"github.com/mfigurski80/NTPeek/types"
	"golang.org/x/exp/maps"
)

func Assign(rows []types.NotionEntry, config PriorityConfig) []Priority {
	tags := make([][]string, len(rows))
	for i, row := range rows {
		tags[i] = make([]string, 0)
		f, ok := row[config.Field].(map[string]interface{})
		if !ok {
			panic(fmt.Errorf(errType.FieldNotFound, config.Field, maps.Keys(row)))
		}
		tagResults, ok := f["multi_select"].([]interface{})
		if !ok {
			panic(fmt.Errorf(errType.UnsupportedType, config.Field))
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
