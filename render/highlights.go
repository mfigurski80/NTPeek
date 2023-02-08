package render

type TextHighlight struct {
	Bg string
	Fg string
}

var colorMap = map[string]TextHighlight{
	"pink":    {"212", "0"},
	"red":     {"203", "0"},
	"orange":  {"#FF9F5E", "0"},
	"yellow":  {"#FFD45E", "0"},
	"green":   {"#A9FF5E", "0"},
	"blue":    {"#5ED4FF", "0"},
	"purple":  {"#9F5EFF", "0"},
	"brown":   {"101", "15"},
	"gray":    {"248", "0"},
	"default": {"240", "15"},
}
