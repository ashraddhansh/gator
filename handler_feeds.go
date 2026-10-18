package main

import (
	"context"
	"fmt"
)


func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching the feeds list: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s, URL: %s, Username: %s\n", feed.FeedName, feed.FeedUrl, feed.UserName)
	}
	return nil
}
