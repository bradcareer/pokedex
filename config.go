package main

var commands = map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
    "help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
    "map": {
        name:        "map",
        description: "Displays the next 20 locations",
        callback:    commandMap,
    },
    "mapb": {
        name:        "mapb",
        description: "Displays the previous 20 locations",
        callback:    commandMapb,
    },}

