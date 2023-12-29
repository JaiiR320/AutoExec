package main

import (
	"autoexec/cli"
	"fmt"
	"os"
)

func expectArg(args []string) {
	if len(args) == 2 {
		fmt.Println("Missing argument")
		os.Exit(1)
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		cli.PrintHelp()
		cli.RunAutoExec()
		return
	}

	switch args[1] {
	case "list":
		cli.AvaliableBuckets()
	case "create":
		expectArg(args)
		cli.CreateBucket(args[2])
	case "add":
		expectArg(args)
		cli.AddPath(args[2])
	case "read":
		expectArg(args)
		cli.ReadBucket(args[2])
	case "run":
		expectArg(args)
		cli.RunBucket(args[2])
	case "help":
		cli.PrintHelp()
	case "startup":
		cli.RunAutoExec()
	default:
		cli.PrintHelp()
		return
	}
}
