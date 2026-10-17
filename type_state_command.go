package main

import (
	"github.com/ashraddhansh/gator/internal/config"
	"github.com/ashraddhansh/gator/internal/database"
)

type state struct {
	db	*database.Queries
	config	*config.Config
}

type command struct {
	name	string
	args	[]string
}

