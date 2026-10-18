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


func handlerFollow(s *state, cmd command, user database.User) error {
	
	if len(cmd.args) !=1 {
		return fmt.Errorf("Please provide the URL to follow")
	}
	feedURL := cmd.args[0]


	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("feed with URL %q not found", feedURL)
		}
		return fmt.Errorf("getting feed: %w", err)
	}

	now := time.Now()

	record, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Error creating feed follow: %w", err)
	}
	fmt.Printf("%s is now following %s\n", record.UserName, record.FeedName)
	return nil
}
