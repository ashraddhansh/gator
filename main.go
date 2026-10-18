package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ashraddhansh/gator/internal/config"
	"github.com/ashraddhansh/gator/internal/database"
	_ "github.com/lib/pq"
)


func main() {
    cfg, err := config.Read()
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

	db, err := sql.Open("postgres", cfg.DbURL)

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

    s := &state{
        config: &cfg,
		db: dbQueries,
    }

    cmds := commands{
        handlers: make(map[string]func(*state, command) error),
    }

    cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)

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
