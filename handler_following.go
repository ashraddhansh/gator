package main

import (
	"context"
	"fmt"

	"github.com/ashraddhansh/gator/internal/database"
)


func handlerFollowing(s *state, cmd command, user database.User) error {


	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error fetching the feeds list for the user: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}
	return nil
}
