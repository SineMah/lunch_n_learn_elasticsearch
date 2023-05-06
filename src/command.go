//go:build command
// +build command

package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"mincedmind.com/elasticsearch/commands/elasticsearch/build"
	"os"
)

func getCommand() (string, []string) {
	commandList := os.Args[1:]

	if len(commandList) == 0 {
		fmt.Println("No command defined")
		os.Exit(1)
	}

	return commandList[0], commandList[1:]
}

func handleCommandValue(value interface{}, args []string) {

	switch value {
	case "build":
		build.Start(args)
	default:
		fmt.Printf("command %s unknown\n", value)
	}
}

func main() {
	godotenv.Load()

	command, args := getCommand()

	handleCommandValue(command, args)
}
