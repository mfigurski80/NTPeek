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
	if len(os.Args) < 2 {
		printTasks(queryNotionTaskDB(NotionDatabaseId))
		return
	}
	// switch command
	markDoneCommand := flag.NewFlagSet("d", flag.ExitOnError)
	switch os.Args[1] {
	case "h", "-h", "--help":
		showUsage()
		return
	case "v", "-v", "--version":
		fmt.Println("nt version:", Version)
		return
	case "d":
		requireDatabaseId()
		markDoneCommand.Parse(os.Args[3:])
		if markDoneCommand.NArg() < 1 {
			fmt.Println("Please provide at least one task ID")
			os.Exit(1)
		}
		markNotionTasksDone(markDoneCommand.Args())
	default:
		fmt.Println("nt: unknown command", os.Args[2])
		fmt.Println()
		showUsage()
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
	fmt.Println("  default -- show tasks from db")
}
