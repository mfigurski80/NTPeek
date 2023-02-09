package render

import (
	"flag"
)

/// `--select` flag provides render string for each task

type SelectRenderString = string

func SetupSelectFlag(fs []*flag.FlagSet) *SelectRenderString {
	var selectFlag string
	for _, f := range fs {
		f.StringVar(&selectFlag, "select", "%Class.R% │ %Name% %Due.relative%", "select render string")
	}
	return &selectFlag
}
