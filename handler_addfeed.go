package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ashraddhansh/gator/internal/database"
	"github.com/google/uuid"
)


func handlerAddFeed(s *state, cmd command) error {

	if len(cmd.args) < 2 {
		return errors.New("Please provide both name and url of the feed")
		os.Exit(1)

	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("No user found with %q name", s.config.CurrentUserName)
			os.Exit(1)
		}
		return err
	}


	feed, err := s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: user.ID,
	})

	if err != nil {
		return fmt.Errorf("Failed to add feed: %w", err)
		os.Exit(1)
	}
	fmt.Printf("Feed created successfully: +%v\n", feed)

	return nil
}
