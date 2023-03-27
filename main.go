package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mfigurski80/NTPeek/filter"
	"github.com/mfigurski80/NTPeek/priority"
	"github.com/mfigurski80/NTPeek/query"
	"github.com/mfigurski80/NTPeek/render"
)

//go:generate bash build/get_auth_token.sh
//go:embed build/auth_token.txt
var defaultAuthorizationSecret string

//go:generate bash build/get_db_id.sh
//go:embed build/db_id.txt
var defaultDBId string
var AccessArgument = query.QueryAccessArgument{
	Secret: defaultAuthorizationSecret,
	DBId:   defaultDBId,
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
		os.Exit(1)
	}

	// parse access
	parseAccessArgument()

	// setup command flag sets
	peekArguments := flag.NewFlagSet("p", flag.ExitOnError)
	allFlagSets := []*flag.FlagSet{peekArguments}

	// setup global flags
	selectRenderString := render.SetupSelectFlag(allFlagSets)
	sortString := query.SetupSortFlag(allFlagSets)
	limitArg := query.SetupLimitFlag(allFlagSets)
	filterStrings := filter.SetupFilterFlag(allFlagSets)
	buildPriorityConfig := priority.SetupPriorityFlags(allFlagSets)

	// check if just peeking
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-") || os.Args[1] == "h" {
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

		params := query.QueryParamArgument{
			Sort:   *sortString,
			Limit:  *limitArg,
			Filter: *filterStrings,
		}
		res, err := query.QueryNotionEntryDB(AccessArgument, params)
		exitOnError(err)

		priorityConfig := buildPriorityConfig()
		fin, err := render.RenderTasks(res, *selectRenderString, priorityConfig)
		fmt.Print(fin)
		exitOnError(err)
	default:
		fmt.Println("nt: unknown command", os.Args)
		fmt.Println()
		showUsage()
	}
}

func parseAccessArgument() {
	// try parse auth secret, db id, in that order
	if len(os.Args[1]) == 50 && strings.HasPrefix(os.Args[1], "secret_") {
		AccessArgument.Secret = os.Args[1]
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}
	if len(os.Args) > 1 && len(os.Args[1]) == 32 {
		AccessArgument.DBId = os.Args[1]
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}
}

func requireAccess(a query.QueryAccessArgument) {
	if a.Secret == "" || a.DBId == "" {
		exitOnError(fmt.Errorf("Please specify a valid Notion Secret Token and Database ID as the first and second positional arguments for this command"))
	}
}

func showUsage() {
	fmt.Printf("Usage: nt [nt-secret?] [nt-database?] <command? [args]>\n\n")
	fmt.Println("Commands:")
	fmt.Println("  v -- show version")
	fmt.Println("  h -- show this help")
	fmt.Println("  p|[none] -- show tasks from db")
}

func exitOnError(err error) {
	if err == nil {
		return
	}
	fmt.Println(
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")).
			Render(err.Error()),
	)
	os.Exit(1)
}
