/*
Copyright (C) 2025  Marius Becker

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	rcon "github.com/butt4cak3/theislercon"
	"github.com/chzyer/readline"
)

var ErrPlayerNotFound = errors.New("player not found")

func main() {
	quiet := false

	serverAddress := os.Getenv("PTEROPROMPT_RCON_ADDRESS")
	rconPassword := os.Getenv("PTEROPROMPT_RCON_PASSWORD")

	argID := 0
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-q":
			quiet = true
		case "-h":
			printHelp(os.Args[0])
			return
		case "--help":
			printHelp(os.Args[0])
			return
		default:
			switch argID {
			case 0:
				serverAddress = arg
			case 1:
				rconPassword = arg
			default:
				printHelp(os.Args[0])
				os.Exit(1)
			}
			argID += 1
		}
	}

	var err error

	for serverAddress == "" {
		serverAddress, err = readline.Line("Server address: ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		serverAddress = strings.TrimSpace(serverAddress)
	}

	for rconPassword == "" {
		pwBytes, err := readline.Password("RCON password: ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		rconPassword = strings.TrimSpace(string(pwBytes))
	}

	client, err := rcon.Connect(serverAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot connect to %s: %v\n", serverAddress, err)
		os.Exit(1)
	}
	defer client.Close()

	err = client.Auth(rconPassword)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot authenticate with %s: %v\n", serverAddress, err)
		os.Exit(1)
	}

	if !quiet {
		fmt.Printf("Connected to %s. Type \"help\" to get a list of available commands.\n", serverAddress)
	}

	rl, err := readline.New("> ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer rl.Close()

Repl:
	for {
		line, err := rl.Readline()
		if err != nil {
			if err == io.EOF {
				break Repl
			}
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		line = strings.TrimSpace(line)

		parts := strings.Split(line, " ")
		if len(parts) == 0 {
			continue
		}

		command := strings.ToLower(parts[0])
		args := parts[1:]

		switch command {
		case "help":
			err = helpCommand(args)
		case "status":
			err = statusCommand(client)
		case "announce":
			err = announceCommand(client, args)
		case "players":
			err = playerListCommand(client)
		case "dm":
			err = messageCommand(client, args)
		case "info":
			err = infoCommand(client, args)
		case "classes":
			err = classesCommand(client, args)
		case "whitelist":
			err = whitelistCommand(client, args)
		case "kick":
			err = kickCommand(client, args)
		case "wipe_corpses":
			err = wipeCorpsesCommand(client)
		case "toggle_gc":
			err = toggleGlobalChatCommand(client)
		case "toggle_humans":
			err = toggleHumansCommand(client)
		case "ai":
			err = aiCommand(client, args)
		case "send":
			err = customCommand(client, args)
		case "quit":
			break Repl
		default:
			fmt.Printf("Unknown command %s. Type \"help\" for a list of commands.\n", command)
			continue
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s command failed: %v\n", command, err)
			os.Exit(1)
		}
	}
}

func printHelp(programName string) {
	fmt.Printf("Usage: %s [-h] [-q] [ ADDRESS [PASSWORD] ]\n", programName)
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("    -h  Show this message")
	fmt.Println("    -q  Print only command outputs")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println("    ADDRESS   Server address and port (optional)")
	fmt.Println("    PASSWORD  RCON password (optional)")
}

// ResolvePlayerName turns a name into an ID.
func ResolvePlayerName(client *rcon.Client, playerName string) (string, error) {
	players, err := client.GetPlayerList()
	if err != nil {
		return "", err
	}

	for _, player := range players {
		if player.Name == playerName {
			return player.ID, nil
		}
	}

	return "", ErrPlayerNotFound
}
