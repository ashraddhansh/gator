package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ashraddhansh/gator/internal/database"
)


func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) { 
				return fmt.Errorf("current user %q does not exist", s.config.CurrentUserName)
				
			}
			return fmt.Errorf("getting current user: %w", err)
		}
		return handler(s, cmd, user)
	}

}


