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

import "fmt"

func helpCommand(args []string) error {
	if len(args) == 0 {
		fmt.Println("Available commands:")
		fmt.Println("    help           Show a list of all commands or details for a specific command")
		fmt.Println("    status         Show some information about the server")
		fmt.Println("    announce       Send a message to all connected players")
		fmt.Println("    players        Show a list of all connected players")
		fmt.Println("    dm             Send a direct message to a specific player")
		fmt.Println("    info           Show detailed information about a specific player")
		fmt.Println("    classes        Manages the list of allowed classes")
		fmt.Println("    whitelist      Manages the whitelist")
		fmt.Println("    kick           Kicks a player from the server")
		fmt.Println("    wipe_corspes   Removes all corpses from the map")
		fmt.Println("    toggle_gc      Toggles the global chat")
		fmt.Println("    toggle_humans  Toggles the humans feature")
		fmt.Println("    ai             Manages AI spawning")
		fmt.Println("    send           Send custom commands")
		fmt.Println("    quit           Exit the program")
		fmt.Println()
		fmt.Println("You can type \"help COMMAND\" to get more information about a specific command.")
		fmt.Println("For example, if you want to know more about the announce command, type \"help announce\".")
	} else {
		switch args[0] {
		case "help":
			fmt.Println("The help command shows a list of all commands or, if passed as an argument, a more detailed explanation of a command.")
			fmt.Println()
			fmt.Println("Usage: help [COMMAND]")
			fmt.Println()
			fmt.Println("Arguments:")
			fmt.Println("    COMMAND  The name of a command")
			fmt.Println()
			fmt.Println("Example: Show detailed help for the dm command")
			fmt.Println("    help dm")
		case "status":
			fmt.Println("The status command shows some information about the server, like the number of currently connected players.")
			fmt.Println()
			fmt.Println("Usage: status")
		case "announce":
			fmt.Println("The announce command sends an announcement message to all players on the server. The message will pop up as a big text box at the top of the screen.")
			fmt.Println()
			fmt.Println("Usage: announce MESSAGE")
			fmt.Println()
			fmt.Println("Arguments:")
			fmt.Println("    MESSAGE  The text you want to send to all players")
			fmt.Println()
			fmt.Println("Example: Announce a server restart")
			fmt.Println("    announce The server will restart in 10 minutes!")
		case "players":
			fmt.Println("The players command shows a list of all currently connected users.")
			fmt.Println()
			fmt.Println("Usage: players")
		case "dm":
			fmt.Println("The dm command sends a direct message to a single player.")
			fmt.Println()
			fmt.Println("Usage: dm PLAYER_NAME MESSAGE")
			fmt.Println()
			fmt.Println("Arguments")
			fmt.Println("    PLAYER_NAME  The name of the recipient")
			fmt.Println("    MESSAGE      The message you want to send")
			fmt.Println()
			fmt.Println("Example: Greet a player")
			fmt.Println("    dm PlayerNameHere Hello!")
		case "info":
			fmt.Println("The info command shows all available information about a specific player, like class, health and position.")
			fmt.Println()
			fmt.Println("Usage: info PLAYER_NAME")
			fmt.Println()
			fmt.Println("Arguments:")
			fmt.Println("    PLAYER_NAME  The name of a player")
			fmt.Println()
			fmt.Println("Example: Get information on the player \"PlayerNameHere\"")
			fmt.Println("    info PlayerNameHere")
		case "classes":
			fmt.Println("The classes command can do several things regarding the list of allowed classes on the server.")
			fmt.Println()
			fmt.Println("Usage: classes SUBCOMMAND [ARGUMENT...]")
			fmt.Println()
			fmt.Println("Subcommands:")
			fmt.Println("    list   Shows a list of all available classes")
			fmt.Println("    allow  Defines which classes are allowed. You have to provide a space-separated list.")
			fmt.Println()
			fmt.Println("Example: Allow only hypsilophodons")
			fmt.Println("    classes allow Hypsilophodon")
		case "whitelist":
			fmt.Println("The whitelist command lets you manage the whitelist on the server.")
			fmt.Println()
			fmt.Println("USAGE: classes SUBCOMMAND [ARGUMENT...]")
			fmt.Println()
			fmt.Println("Subcommands:")
			fmt.Println("    status  Shows whether the whitelist is currently turned on or off")
			fmt.Println("    toggle  Turns the whitelist on or off")
			fmt.Println("    add     Adds one or more players to the whitelist")
			fmt.Println("    remove  Removes one or more players from the whitelist")
			fmt.Println()
			fmt.Println("The add and remove commands will try to resolve player names to IDs for you. If the player that you want to add/remove is not currently playing on the server, you have to use the ID directly.")
			fmt.Println()
			fmt.Println("Example: Add two players to the whitelist")
			fmt.Println("    whitelist add FirstPlayer SecondPlayer")
		case "kick":
			fmt.Println("The kick command kicks a currently connected player from the server. You can provide a message that will be shown to the player in the menu.")
			fmt.Println("")
			fmt.Println("Usage: kick PLAYER_NAME [REASON]")
			fmt.Println()
			fmt.Println("Arguments:")
			fmt.Println("    PLAYER_NAME  Name of the player you want to kick")
			fmt.Println("    REASON       A message that will be shown to the player in the menu")
			fmt.Println()
			fmt.Println("Example: Kick a player")
			fmt.Println("    kick PlayerNameHere You have broken the law")
		case "send":
			fmt.Println("The send command enables you to send commands to the server that this tool doesn't support yet.")
			fmt.Println()
			fmt.Println("Usage: send CODE [ARGUMENT...]")
			fmt.Println()
			fmt.Println("Arguments:")
			fmt.Println("    CODE      The 2-digit hexadecimal code of the message type you want to send")
			fmt.Println("    ARGUMENT  (optional) The arguments that you want to send with your command")
			fmt.Println()
			fmt.Println("Examples: Send an announcement")
			fmt.Println("    send 10 Testing")
		case "wipe_corspes":
			fmt.Println("The wipe_corspes command removes all corpses from the map to improve performance.")
			fmt.Println()
			fmt.Println("Usage: wipe_corpses")
		case "toggle_gc":
			fmt.Println("The toggle_gc turns the global chat on or off.")
			fmt.Println()
			fmt.Println("Usage: toggle_gc")
		case "togle_humans":
			fmt.Println("The toggle_humans command turns the humans feature of the game on or off.")
			fmt.Println()
			fmt.Println("Usage: toggle_humans")
		case "ai":
			fmt.Println("The ai command lets you manage AI spawns.")
			fmt.Println()
			fmt.Println("Usage: ai SUBCOMMAND [ARGUMENT...]")
			fmt.Println()
			fmt.Println("Subcommands:")
			fmt.Println("    list     Shows a list of all AI classes")
			fmt.Println("    toggle   Turns AI spawning on or off")
			fmt.Println("    disable  Disables one or more AI classes. You have to provide a space separated list.")
			fmt.Println("             You can also pass \"all\" or \"none\" to disable all or no AI classes respectively.")
			fmt.Println("    density  Lets you control how much AI spawns. You have to pass a number.")
			fmt.Println()
			fmt.Println("Example: Disable boars")
			fmt.Println("    ai disable Boar")
		case "quit":
			fmt.Println("The quit command exits this program.")
			fmt.Println()
			fmt.Println("Usage: quit")
		default:
			helpCommand([]string{})
		}
	}
	return nil
}
