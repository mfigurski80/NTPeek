package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mfigurski80/NTPeek/priority"
	"github.com/mfigurski80/NTPeek/render"
)

//go:generate bash build/get_auth_token.sh
//go:embed build/auth_token.txt
var NotionAuthorizationSecret string

//go:generate bash build/get_version.sh
//go:embed build/version.txt
var Version string

var NotionDatabaseId string

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
		NotionAuthorizationSecret = os.Args[1]
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}
	if len(os.Args) > 1 && len(os.Args[1]) == 32 {
		NotionDatabaseId = os.Args[1]
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}

	// setup command flag sets
	markDoneArguments := flag.NewFlagSet("d", flag.ExitOnError)
	peekArguments := flag.NewFlagSet("p", flag.ExitOnError)
	allFlagSets := []*flag.FlagSet{markDoneArguments, peekArguments}

	// setup global flags
	applyFn := []func(){
		setupGlobalFieldNameFlags(allFlagSets),
		priority.SetupGlobalTagPriorityFlags(allFlagSets),
	}
	selectRenderString := render.SetupSelectFlag(allFlagSets)

	// check if just peeking
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-") {
		os.Args = append([]string{os.Args[0], "p"}, os.Args[1:]...)
	}

	// parse user command
	switch os.Args[1] {
	case "v", "version":
		fmt.Println("nt version:", Version)
		return
	case "d", "do", "delete", "done":
		requireDatabaseId()
		markDoneArguments.Parse(os.Args[2:])
		if markDoneArguments.NArg() < 1 {
			fmt.Println("Please provide at least one task ID")
			os.Exit(1)
		}
		markNotionTasksDone(markDoneArguments.Args())
	case "p", "peek":
		requireDatabaseId()
		peekArguments.Parse(os.Args[2:])
		for _, fn := range applyFn {
			fn()
		}
		render.RenderTasks(queryNotionEntryDB(NotionDatabaseId), *selectRenderString)
	default:
		fmt.Println("nt: unknown command", os.Args)
		fmt.Println()
		showUsage()
	}
}

func requireDatabaseId() {
	if NotionDatabaseId == "" {
		fmt.Println("Please specify a valid Notion database ID as the first argument for this command")
		os.Exit(1)
	}
}

func showUsage() {
	fmt.Println("Usage: nt [database-id?] <command? [args]>\n")
	fmt.Println("Commands:")
	fmt.Println("  d [task-id] ... -- mark task(s) from db as done")
	fmt.Println("  v -- show version")
	fmt.Println("  h -- show this help")
	fmt.Println("  p|[none] -- show tasks from db")
}
