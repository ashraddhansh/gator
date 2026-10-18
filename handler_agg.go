package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ashraddhansh/gator/internal/rss"
)

func scrapeFeeds(s * state) error {
	fetchedFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("No feeds available to fetch")
		}
		return fmt.Errorf("getting next feed: %w", err)
	}

	feed, err := rss.FetchFeed(context.Background(), fetchedFeed.Url)
	if err != nil {
		return fmt.Errorf("fetching feed %q: %w", fetchedFeed.Name, err)
	}

	err = s.db.MarkFeedFetched(context.Background(), fetchedFeed.ID)

	if err != nil {
		return fmt.Errorf("marking the feed as fetched: %w", err)
	}


	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
	}
	return nil
}


func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) !=1 {
		return fmt.Errorf("usage: agg <time_between_reqs>")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Parsing time between requests: %w", err)
	}
	if timeBetweenReqs <= 0 {
		return fmt.Errorf("time between requests must be greater than zero")
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	defer ticker.Stop()

	for {
		if err := scrapeFeeds(s); err != nil {
			fmt.Printf("Error scraping feeds: %v\n", err)
		}

		<- ticker.C
	}
}
