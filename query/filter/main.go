package filter

import (
	"github.com/alecthomas/participle/v2"
)

func ParseFilter(filter string) string {
	// fmt.Println("Got Filter: ", filter)
	p := setupParser()
	f, err := p.ParseString("", "("+filter+")")
	if err != nil {
		panic(err)
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
		panic(err)
	}
	return parser
}
