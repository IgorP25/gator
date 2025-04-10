package main

import (
	"log"
	"os"

	"github.com/IgorP25/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	programState := &state{
		cfg: &cfg,
	}

	commands := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}

	commands.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	err = commands.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
