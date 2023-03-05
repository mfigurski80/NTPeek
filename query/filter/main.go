package filter

import (
	"github.com/alecthomas/participle/v2"
)

func ParseFilter(filter string) (string, error) {
	// fmt.Println("Got Filter: ", filter)
	p := setupParser()
	f, err := p.ParseString("", "("+filter+")")
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
