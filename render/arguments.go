package render

import (
	"flag"
)

/// `--select` flag provides render string for each task

type SelectRenderString = string

func SetupSelectFlag(fs []*flag.FlagSet) *SelectRenderString {
	var selectFlag string
	for _, f := range fs {
		f.StringVar(&selectFlag, "select", "%Class:right% │%_p%%Name% %Due:relative%", "select render string, describing which fields to print (with optional modifiers)")
	}
	return &selectFlag
}
