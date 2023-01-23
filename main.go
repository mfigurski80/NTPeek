package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

var NotionDatabaseId string

//go:generate bash build/get_auth_token.sh
//go:embed build/auth_token.txt
var NotionAuthorizationSecret string

//go:generate bash build/get_version.sh
//go:embed build/version.txt
var Version string

var FieldNamesConfig FieldNames

func main() {

	// insufficient args
	if len(os.Args) < 2 {
		fmt.Println("nt: insufficient arguments", os.Args)
		fmt.Println()
		showUsage()
		return
	}

	// parse db id
	if len(os.Args[1]) == 32 {
		NotionDatabaseId = os.Args[1]
		os.Args = os.Args[1:]
	}

	// parse generic commands
	FieldNamesConfig = parseFieldNameArguments(os.Args[1:])

	// switch user action
	if len(os.Args) < 2 {
		printTasks(queryNotionTaskDB(NotionDatabaseId))
		return
	}
	markDoneArguments := flag.NewFlagSet("d", flag.ExitOnError)
	switch os.Args[1] {
	case "h", "-h", "--help":
		showUsage()
		return
	case "v", "-v", "--version":
		fmt.Println("nt version:", Version)
		return
	case "d":
		requireDatabaseId()
		markDoneArguments.Parse(os.Args[2:])
		if markDoneArguments.NArg() < 1 {
			fmt.Println("Please provide at least one task ID")
			os.Exit(1)
		}
		markNotionTasksDone(markDoneArguments.Args())
	default:
		printTasks(queryNotionTaskDB(NotionDatabaseId))
	}
}

func requireDatabaseId() {
	if NotionDatabaseId == "" {
		fmt.Println("Please specify a Notion database ID")
		os.Exit(1)
	}
}

func showUsage() {
	fmt.Println("Usage: nt [database-id?] <command? [args]>\n")
	fmt.Println("Commands:")
	fmt.Println("  d [task-id] ... -- mark task(s) from db as done")
	fmt.Println("  v -- show version")
	fmt.Println("  h -- show this help")
	fmt.Println("  [default] -- show tasks from db")
}
