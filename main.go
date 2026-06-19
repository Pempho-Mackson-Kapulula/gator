package main

import (
	"log"
	"os"

	"github.com/Pempho-Mackson-Kapulula/gator/internal/config"
)

// state struct
type state struct {
	cfg *config.Config
}

func main() {

	//read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	stateinstance := &state{
		cfg: &cfg,
	}

	commandsInstance := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	commandsInstance.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("")
	}

	cmdName := os.Args[1]
	args := os.Args[2:]

	command := command{
		name: cmdName,
		args: args,
	}

	err = commandsInstance.run(stateinstance, command)
	if err != nil {
		log.Fatal(err)
	}

}
