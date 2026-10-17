package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Please provide the username")

	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])

	if errors.Is(err, sql.ErrNoRows) {
    	return fmt.Errorf("user %q does not exist", cmd.args[0])
		os.Exit(1)
	}
	if err != nil {
    	return err
	}


	if err := s.config.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Println("The user has been set")
	return nil
}
