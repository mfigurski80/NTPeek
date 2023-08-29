package renderField

import "github.com/mfigurski80/NTPeek/priority"

type RenderRowConfig struct {
	Name      string
	Modifiers []string
	Priority  *[]priority.Priority
}

type RenderRowFunction func([]interface{}, RenderRowConfig) ([]string, error)
