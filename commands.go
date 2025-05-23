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
	"net"
	"strconv"
	"strings"

	rcon "github.com/butt4cak3/theislercon"
	"golang.org/x/text/message"
)

func statusCommand(client *rcon.Client) error {
	details, err := client.GetServerDetails()
	if err != nil {
		return err
	}

	yesno := func(v bool) string {
		if v {
			return "yes"
		} else {
			return "no"
		}
	}

	fmt.Println("Server information:")
	fmt.Printf("    Name:                %s\n", details.Name)
	fmt.Printf("    Players:             %d/%d\n", details.CurrentPlayers, details.MaxPlayers)
	fmt.Printf("    Day/Night length:    %d/%d minutes\n", details.DayLengthMinutes, details.NightLengthMinutes)
	fmt.Printf("    Password protected?  %s\n", yesno(details.HasPassword))
	fmt.Printf("    Global chat enabled? %s\n", yesno(details.EnableGlobalChat))
	fmt.Printf("    Queue enabled?       %s\n", yesno(details.QueueEnabled))
	fmt.Printf("    Whitelist enabled?   %s\n", yesno(details.Whitelist))

	return nil
}

func announceCommand(client *rcon.Client, args []string) error {
	if len(args) < 1 {
		fmt.Println("Missing MESSAGE")
		return nil
	}

	message := strings.Join(args, " ")
	return client.Announce(message)
}

func playerListCommand(client *rcon.Client) error {
	players, err := client.GetPlayerList()
	if err != nil {
		return err
	}

	fmt.Println("Connected players:")

	for _, player := range players {
		fmt.Printf("    %s\n", player.Name)
	}

	return nil
}

func messageCommand(client *rcon.Client, args []string) error {
	if len(args) < 1 {
		fmt.Println("Missing PLAYER_NAME")
		return nil
	} else if len(args) < 2 {
		fmt.Println("Missing MESSAGE")
		return nil
	}

	playerID, err := ResolvePlayerName(client, args[0])
	if err != nil {
		if errors.Is(err, ErrPlayerNotFound) {
			fmt.Printf("Player \"%s\" not found\n", args[0])
			return nil
		}
		return err
	}
	message := strings.Join(args[1:], " ")
	err = client.SendDirectMessage(playerID, message)
	if err != nil {
		return err
	}

	return nil
}

func infoCommand(client *rcon.Client, args []string) error {
	if len(args) < 1 {
		fmt.Println("Missing PLAYER_NAME")
		return nil
	}

	playerName := strings.ToLower(args[0])

	players, err := client.GetPlayerData()
	if err != nil {
		return err
	}

	for _, player := range players {
		if strings.ToLower(player.Name) == playerName {
			p := message.NewPrinter(message.MatchLanguage("en"))
			fmt.Printf("Player %s\n", player.Name)
			fmt.Printf("    ID:       %s\n", player.ID)
			fmt.Printf("    Class:    %s\n", player.DinoClass.Name())
			fmt.Printf("    Growth:   %d%%, Health: %d%%, Stamina: %d%%, Hunger: %d%%, Thirst: %d%%\n", player.Growth, player.Health, player.Stamina, player.Hunger, player.Thirst)
			p.Printf("    Location: %.3f, %.3f, %.3f\n", player.Location.Y, player.Location.X, player.Location.Z)
			return nil
		}
	}

	fmt.Printf("Player \"%s\" not found\n", args[0])

	return nil
}

func classesCommand(client *rcon.Client, args []string) error {
	if len(args) == 0 {
		fmt.Println("No subcommand provided")
		fmt.Println("Type \"help classes\" to learn more about this command.")
		return nil
	}

	cmd, args := strings.ToLower(args[0]), args[1:]

	switch cmd {
	case "list":
		fmt.Println("List of all classes:")
		for _, class := range rcon.AllClasses {
			fmt.Printf("    %s\n", class)
		}
		return nil
	case "allow":
		var classes []rcon.DinoClass
		if len(args) == 0 {
			fmt.Println("No classes provided.")
			fmt.Println("Type \"classes list\" to get a list of all available classes or \"classes allow all\" to allow all classes at the same time.")
			return nil
		}
		if len(args) == 1 && args[0] == "all" {
			classes = rcon.AllClasses[:]
		} else {
			classes = make([]rcon.DinoClass, len(args))
			for i, arg := range args {
				if !rcon.IsClass(arg) {
					fmt.Printf("\"%s\" is not a class. Type \"classes list\" to get a list of all classes.", arg)
					return nil
				}
				classes[i] = rcon.DinoClass(arg)
			}
		}
		return client.UpdatePlayables(classes)
	default:
		fmt.Printf("Invalid subcommand \"%s\".\n", cmd)
		fmt.Println("Type \"help classes\" to learn more about this command.")
		return nil
	}
}

func whitelistCommand(client *rcon.Client, args []string) error {
	if len(args) == 0 {
		fmt.Println("No subcommand provided.")
		fmt.Println("Type \"help whitelist\" to learn more about this command.")
		return nil
	}

	cmd, args := strings.ToLower(args[0]), args[1:]

	switch cmd {
	case "toggle":
		status, err := client.ToggleWhitelist()
		if err != nil {
			return err
		}
		if status {
			fmt.Println("The whitelist is now on")
		} else {
			fmt.Println("The whitelist is nof off")
		}
		return nil
	case "status":
		details, err := client.GetServerDetails()
		if err != nil {
			return err
		}
		if details.Whitelist {
			fmt.Println("The whitelist is currently on")
		} else {
			fmt.Println("The whitelist is currently off")
		}
		return nil
	case "add":
		if len(args) == 0 {
			fmt.Println("No PlayerIDs provided.")
			return nil
		}
		playerIDs := make([]string, len(args))
		for i, name := range args {
			id, err := ResolvePlayerName(client, name)
			if err != nil {
				if err == ErrPlayerNotFound {
					playerIDs[i] = name
					continue
				}
				return err
			}
			playerIDs[i] = id
		}
		err := client.AddWhitelistID(playerIDs...)
		if err != nil {
			return err
		}
		fmt.Printf("Added %d IDs to the whitelist: %s\n", len(playerIDs), strings.Join(playerIDs, ", "))
		return nil
	case "remove":
		if len(args) == 0 {
			fmt.Println("No PlayerIDs provided.")
			return nil
		}
		playerIDs := make([]string, len(args))
		for i, name := range args {
			id, err := ResolvePlayerName(client, name)
			if err != nil {
				if err == ErrPlayerNotFound {
					playerIDs[i] = name
					continue
				}
				return err
			}
			playerIDs[i] = id
		}
		err := client.RemoveWhitelistID(playerIDs...)
		if err != nil {
			return err
		}
		fmt.Printf("Removed %d IDs from the whitelist: %s\n", len(playerIDs), strings.Join(playerIDs, ", "))
		return nil
	default:
		fmt.Printf("Invalid subcommand \"%s\".\n", cmd)
		fmt.Println("Type \"help whitelist\" to learn more about this command.")
		return nil
	}
}

func kickCommand(client *rcon.Client, args []string) error {
	if len(args) == 0 {
		fmt.Println("Missing player name")
		return nil
	}

	playerName := args[0]
	var reason string
	if len(args) > 1 {
		reason = strings.Join(args[1:], " ")
	} else {
		reason = "You were kicked from the server."
	}
	playerID, err := ResolvePlayerName(client, playerName)
	if err != nil {
		if err == ErrPlayerNotFound {
			fmt.Printf("No such player \"%s\"\n", playerName)
			return nil
		}
		return err
	}
	err = client.KickPlayer(playerID, reason)
	if err != nil {
		return err
	}
	fmt.Printf("%s was kicked from the server. Reason: %s\n", playerName, reason)
	return nil
}

func wipeCorpsesCommand(client *rcon.Client) error {
	err := client.WipeCorpses()
	if err != nil {
		return err
	}
	fmt.Println("Corpses wiped")
	return nil
}

func toggleGlobalChatCommand(client *rcon.Client) error {
	state, err := client.ToggleGlobalChat()
	if err != nil {
		return err
	}
	if state {
		fmt.Println("Global chat is now on")
	} else {
		fmt.Println("Global chat is now off")
	}
	return err
}

func toggleHumansCommand(client *rcon.Client) error {
	status, err := client.ToggleHumans()
	if err != nil {
		return err
	}
	if status {
		fmt.Println("Humans are now on")
	} else {
		fmt.Println("Humans are now off")
	}
	return err
}

func aiCommand(client *rcon.Client, args []string) error {
	if len(args) == 0 {
		fmt.Println("No subcommand provided.")
		fmt.Println("Type \"help whitelist\" to learn more about this command.")
		return nil
	}

	cmd, args := strings.ToLower(args[0]), args[1:]

	switch cmd {
	case "list":
		fmt.Println("List of all AI classes:")
		for _, class := range rcon.AllAIClasses {
			fmt.Printf("    %s\n", class)
		}
		return nil
	case "toggle":
		status, err := client.ToggleAI()
		if err != nil {
			return err
		}
		if status {
			fmt.Println("AI spawns are now on")
		} else {
			fmt.Println("AI spawns are now off")
		}
		return nil
	case "disable":
		var classes []rcon.AIClass
		if len(args) == 0 {
			fmt.Println("No classes provided.")
			fmt.Println("Type \"ai list\" to get a list of all available classes or \"classes disable all|none\" to disable all or no classes at all.")
			return nil
		}
		if len(args) == 1 && args[0] == "all" {
			classes = rcon.AllAIClasses[:]
		} else if len(args) == 1 && args[0] == "none" {
			// Keep empty
		} else {
			classes = make([]rcon.AIClass, len(args))
			for i, arg := range args {
				if !rcon.IsAIClass(arg) {
					fmt.Printf("\"%s\" is not an AI class. Type \"classes list to get a list of all classes.\n", arg)
					return nil
				}
				classes[i] = rcon.AIClass(arg)
			}
		}
		err := client.DisableAIClasses(classes)
		if err != nil {
			return err
		}
		fmt.Println("Updated AI classes.")
		return nil
	case "density":
		if len(args) == 0 {
			fmt.Println("No density provided.")
			return nil
		}
		density, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			fmt.Println("The density must be a number.")
			return nil
		}
		err = client.SetAIDensity(float32(density))
		if err != nil {
			return err
		}
		fmt.Println("Updated AI density.")
		return nil
	default:
		fmt.Printf("Invalid subcommand \"%s\".\n", cmd)
		fmt.Println("Type \"help ai\" to learn more about this command.")
		return nil
	}
}

func customCommand(client *rcon.Client, args []string) error {
	if len(args) < 1 {
		fmt.Println("Missing command byte")
		return nil
	}

	commandByte, err := strconv.ParseUint(args[0], 16, 8)
	if err != nil {
		fmt.Println("Command byte must be a hexadecimal number, e.g. 3a")
		return nil
	}

	response, err := client.ExecCommand(byte(commandByte), args[1:]...)
	if err != nil {
		if ne, ok := err.(net.Error); ok && ne.Timeout() {
			fmt.Println("The server did not respond with anything.")
			return nil
		}
		return err
	}

	fmt.Println(response)
	return nil
}
