package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ashraddhansh/gator/internal/database"
)


func handlerUnfollow(s *state, cmd command, user database.User) error {
	
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
	err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("unfollowing feed: %w", err)
	}
	fmt.Printf("Feed unfollowed successfully: %s", feedURL)

	return nil
}
