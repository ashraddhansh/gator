package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Please provide the username")

	}

	if err := s.config.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Println("The user has been set")
	return nil
}
