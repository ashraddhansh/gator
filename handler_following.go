package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)


func handlerFollowing(s *state, cmd command) error {

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)

	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("user %q does not exist", s.config.CurrentUserName)
	}
	if err != nil {
    	return err
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error fetching the feeds list for the user: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}
	return nil
}
