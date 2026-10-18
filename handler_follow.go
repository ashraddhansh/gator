package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ashraddhansh/gator/internal/database"
	"github.com/google/uuid"
)


func handlerFollow(s *state, cmd command) error {
	
	if len(cmd.args) !=1 {
		return fmt.Errorf("Please provide the URL to follow")
	}
	feedURL := cmd.args[0]


	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)

	if err != nil {
		return fmt.Errorf("getting current user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("feed with URL %q not found", feedURL)
		}
		return fmt.Errorf("getting feed: %w", err)
	}

	record, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Error creating feed follow: %v", err)
	}
	fmt.Printf("%s is now following %s\n", record.UserName, record.FeedName)
	return nil
}
