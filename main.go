package main

import (
	"time"

	"https://github.com/bradcareer/pokedex"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
