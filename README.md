# The Isle RCON Console

This is a terminal application (i.e. no graphical interface) that lets you send commands to your The Isle Evrima server.

## Connecting to your server

To run the program, simply run `./theisleconsole` on Linux or `.\theisleconsole.exe` on Windows. There will be a prompt that asks you for your server's RCON address and password.

Alternatively, you can pass the address, or both address and password as command line parameters like so:

```sh
./theisleconsole 127.0.0.1:8888 YourSecurePasswordHere
```

As a third option, you can set the environment variables `THEISLE_RCON_ADDRESS` and/or `THEISLE_RCON_PASSWORD` before you start the program and it will use those values instead.

## Usage

Once you're connected to your server, you will see a prompt (`>`). From here, you can type commands and send them with enter. If you're unsure what commands are available or how to use them, you can type `help` to get a list of commands, or `help COMMAND` to get more information about one specific command.

### List of commands

| Command       | Description                                                   |
| ------------- | ------------------------------------------------------------- |
| help          | Show a list of all commands or details for a specific command |
| status        | Show some information about the server                        |
| announce      | Send a message to all connected players                       |
| players       | Show a list of all connected players                          |
| dm            | Send a direct message to a specific player                    |
| info          | Show detailed information about a specific player             |
| classes       | Manages the list of allowed classes                           |
| whitelist     | Manages the whitelist                                         |
| kick          | Kicks a player from the server                                |
| wipe_corspes  | Removes all corpses from the map                              |
| toggle_gc     | Toggles the global chat                                       |
| toggle_humans | Toggles the humans feature                                    |
| ai            | Manages AI spawning                                           |
| send          | Send custom commands                                          |
| quit          | Exit the program                                              |
