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


func handlerAddFeed(s *state, cmd command) error {

	if len(cmd.args) < 2 {
		return errors.New("Please provide both name and url of the feed")

	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("No user found with %q name", s.config.CurrentUserName)
		}
		return err
	}

	now := time.Now()
	feed, err := s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: user.ID,
	})
	if err != nil {
    	return fmt.Errorf("failed to add feed: %w", err)
    }

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("Feed created but failed to follow it: %w", err)
	}

	fmt.Printf("Feed created successfully: +%v\n", feed)
	fmt.Printf("%s is now following %s", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
