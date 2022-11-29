package main

import (
	"flag"
	"fmt"
	"os"
)

var NotionDatabaseId = "d048f752003e4c199533c9a39608917e"
var NotionAuthorizationSecret string
var Version string

func main() {
	if len(os.Args) < 2 { // no args
		printTasks(queryNotionTaskDB())
	} else {
		markDoneCommand := flag.NewFlagSet("d", flag.ExitOnError)
		switch os.Args[1] {
		case "d":
			markDoneCommand.Parse(os.Args[2:])
			if markDoneCommand.NArg() < 1 {
				fmt.Println("Please provide at least one task ID")
				os.Exit(1)
			}
			markNotionTasksDone(markDoneCommand.Args())
		case "v":
			fmt.Println(Version)
		default:
			fmt.Println("Unknown command:", os.Args[1])
			flag.PrintDefaults()
		}
	}

}
