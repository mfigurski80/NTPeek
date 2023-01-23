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
	// switch arg 1
	switch os.Args[1] {
	case "h", "-h", "--help":
		showUsage()
		return
	case "v", "-v", "--version":
		fmt.Println("nt version: %s", Version)
		return
	default:
		NotionDatabaseId = os.Args[1]
	}
	// switch arg 2
	if len(os.Args) == 2 {
		printTasks(queryNotionTaskDB(NotionDatabaseId))
		return
	}
	markDoneCommand := flag.NewFlagSet("d", flag.ExitOnError)
	switch os.Args[2] {
	case "d":
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

func showUsage() {
	fmt.Println("Usage: nt [database-id?] <command? [args]>\n")
	fmt.Println("Commands:")
	fmt.Println("  d [task-id] ... -- mark task(s) from db as done")
	fmt.Println("  v -- show version")
	fmt.Println("  h -- show this help")
	fmt.Println("  default -- show tasks from db")
}
