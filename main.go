package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mfigurski80/NTPeek/priority"
	"github.com/mfigurski80/NTPeek/query"
	"github.com/mfigurski80/NTPeek/render"
)

//go:generate bash build/get_auth_token.sh
//go:embed build/auth_token.txt
var notionAuthorizationSecret string
var AccessArgument = query.QueryAccessArgument{
	notionAuthorizationSecret,
	"",
}

//go:generate bash build/get_version.sh
//go:embed build/version.txt
var Version string

func main() {

	// insufficient args
	if len(os.Args) < 2 {
		fmt.Println("nt: insufficient arguments", os.Args)
		fmt.Println()
		showUsage()
		return
	}

	// try parse auth secret, db id
	if len(os.Args[1]) == 50 && strings.HasPrefix(os.Args[1], "secret_") {
		AccessArgument.Secret = os.Args[1]
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}
	if len(os.Args) > 1 && len(os.Args[1]) == 32 {
		AccessArgument.DBId = os.Args[1]
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}

	// setup command flag sets
	peekArguments := flag.NewFlagSet("p", flag.ExitOnError)
	allFlagSets := []*flag.FlagSet{peekArguments}

	// setup global flags
	applyFn := []func(){
		query.SetupFieldNameFlags(allFlagSets),
		priority.SetupGlobalTagPriorityFlags(allFlagSets),
	}
	selectRenderString := render.SetupSelectFlag(allFlagSets)
	sortString := query.SetupSortFlag(allFlagSets)
	filterString := query.SetupFilterFlag(allFlagSets)

	// check if just peeking
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-") {
		os.Args = append([]string{os.Args[0], "p"}, os.Args[1:]...)
	}

	// parse user command
	switch os.Args[1] {
	case "v", "version":
		fmt.Println("nt version:", Version)
		return
	case "p", "peek":
		requireAccess(AccessArgument)
		peekArguments.Parse(os.Args[2:])
		for _, fn := range applyFn {
			fn()
		}
		params := query.QueryParamArgument{*sortString, *filterString}
		res := query.QueryNotionEntryDB(AccessArgument, params)
		render.RenderTasks(res, *selectRenderString)
	default:
		fmt.Println("nt: unknown command", os.Args)
		fmt.Println()
		showUsage()
	}
}

func requireAccess(a query.QueryAccessArgument) {
	if a.Secret == "" || a.DBId == "" {
		fmt.Println("Please specify a valid Notion Secret Token and Database ID as the first and second positional arguments for this command")
		os.Exit(1)
	}
}

func showUsage() {
	fmt.Println("Usage: nt [nt-secret?] [nt-database?] <command? [args]>\n")
	fmt.Println("Commands:")
	fmt.Println("  v -- show version")
	fmt.Println("  h -- show this help")
	fmt.Println("  p|[none] -- show tasks from db")
}
