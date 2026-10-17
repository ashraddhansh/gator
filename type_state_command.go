package main

import "github.com/ashraddhansh/gator/internal/config"

type state struct {
	config	*config.Config
}

type command struct {
	name	string
	args	[]string
}

