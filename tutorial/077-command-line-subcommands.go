package main

import (
	"flag"
	"fmt"
	"os"
)

func commandlineSubcommands() {
	print := fmt.Println

	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	message := commitCmd.String("m", "", "commit message")

	remoteCmd := flag.NewFlagSet("remote", flag.ExitOnError)
	remoteGit := remoteCmd.String("add", "", "add remote git")
	removeGit := remoteCmd.Bool("remove", false, "remove remote git")

	flag.Parse()

	if len(os.Args) < 2 {
		print("expected 'commit' or 'remote' subcommands.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "commit":
		commitCmd.Parse(os.Args[2:])
		print(*message)
	case "remote":
		remoteCmd.Parse(os.Args[2:])
		print(*remoteGit, *removeGit)
	default:
		print("expected 'commit' or 'remote' subcommands.")
		os.Exit(1)
	}
}
