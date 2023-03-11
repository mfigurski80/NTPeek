package filter

import (
	"strings"

	"github.com/alecthomas/participle/v2"
)

func ParseFilter(filter []FilterString) (string, error) {
	// build composite
	if len(filter) == 0 {
		return "", nil
	}
	comp := "(" + filter[0] + ")"
	if len(filter) > 1 {
		comp = "((" + strings.Join(filter, ") AND (") + "))"
	}
	// parse
	p := setupParser()
	f, err := p.ParseString("", comp)
	if err != nil {
		return "", err
	}
	// fmt.Println("Rendered:", f.Render())
	return f.Render()
}

func setupParser() *participle.Parser[condition] {
	parser, err := participle.Build[condition](
		participle.Unquote("String"),
		valueUnion,
	)
	if err != nil {
		panic("UNREACHABLE unable to build parser: " + err.Error())
	}
	return parser
}
