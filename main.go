package main

import (
	"fmt"
	"os"

	"github.com/ashraddhansh/gator/internal/config"
)


func main() {
    cfg, err := config.Read()
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    s := &state{
        config: &cfg,
    }

    cmds := commands{
        handlers: make(map[string]func(*state, command) error),
    }

    cmds.register("login", handlerLogin)

    args := os.Args
    if len(args) < 2 {
        fmt.Println("Not enough arguments")
		os.Exit(1)
    }

    cmd := command{
        name: args[1],
        args: args[2:],
    }

    if err := cmds.run(s, cmd); err != nil {
        fmt.Printf("Error: %v\n", err)
		os.Exit(1)
    }
}
