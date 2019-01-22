package main

import (
	"log"
	"fmt"
	"os"

	"github.com/mit-dci/opencx/cxrpc"
)

// opencx-cli is the client, opencx is the server
func main() {
	commandArg := os.Args[1:]

	// TODO just for now
	err := parseCommands(commandArg)
	if err != nil {
		log.Fatal(err)
	}
}

func parseCommands(commands []string) error {
	var args []string

	if len(commands) == 0 {
		return fmt.Errorf("Please specify arguments for exchange CLI")
	}
	cmd := commands[0]

	client := new(cxrpc.OpencxRPCClient)
	if len(commands) > 1 {
		args = commands[1:]
	}
	if cmd == "register" {
		if len(args) != 2 {
			return fmt.Errorf("Must specify two arguments: username and password. Instead, %d arguments were specified", len(args))
		}

		username := args[0]
		password := args[1]

		// construct JSON and send through rpc
		err := client.Register(username, password)
		if err != nil {
			return fmt.Errorf("Error registering, probably username is taken")
		}
		// method that uses rpc should also set token to instance of client if returned
	}
	if cmd == "login" {
		if len(args) != 2 {
			return fmt.Errorf("Must specify two arguments: username and password. Instead, %d arguments were specified", len(args))
		}
	}
	if cmd == "vieworderbook" {
		// run method that returns orders in json
	}
	return nil
}