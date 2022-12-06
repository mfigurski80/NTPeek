package main

type TextHighlight struct {
	Bg   string
	Fore string
}

var colorMap = map[string]TextHighlight{
	"pink":    {"212", "0"},
	"red":     {"203", "0"},
	"orange":  {"208", "0"},
	"yellow":  {"221", "0"},
	"green":   {"120", "0"},
	"blue":    {"39", "0"},
	"purple":  {"141", "0"},
	"brown":   {"101", "15"},
	"gray":    {"248", "0"},
	"default": {"240", "15"},
}
