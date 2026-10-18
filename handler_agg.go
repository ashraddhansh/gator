package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
	"log"

	"github.com/ashraddhansh/gator/internal/database"
	"github.com/ashraddhansh/gator/internal/rss"
	"github.com/google/uuid"
)

func parsePublishedAt(pubDate string) (sql.NullTime, error) {
	pubDate = strings.TrimSpace(pubDate)


	if pubDate == "" {
		return sql.NullTime{}, nil
	}
	layouts := []string{
    	time.RFC1123Z,
        time.RFC1123,
        time.RFC822Z,
        time.RFC822,
        time.RFC850,
        time.RFC3339,
        time.RFC3339Nano,
        time.ANSIC,
        time.RubyDate,
    }

	for _, layout := range layouts {
		parsed, err := time.Parse(layout, pubDate)
		if err == nil {
			return sql.NullTime{
				Time: parsed,
				Valid: true,
			}, nil
		}
	}

	return sql.NullTime{}, fmt.Errorf("unsupported publication date format: %q", pubDate)

}

func scrapeFeeds(s *state) error {
    ctx := context.Background()

    fetchedFeed, err := s.db.GetNextFeedToFetch(ctx)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return fmt.Errorf("no feeds available to fetch")
        }
        return fmt.Errorf("getting next feed: %w", err)
    }

    feed, err := rss.FetchFeed(ctx, fetchedFeed.Url)
    if err != nil {
        return fmt.Errorf("fetching feed %q: %w", fetchedFeed.Name, err)
    }

    if err := s.db.MarkFeedFetched(ctx, fetchedFeed.ID); err != nil {
        return fmt.Errorf("marking feed as fetched: %w", err)
    }

    for _, item := range feed.Channel.Item {
        publishedAt, err := parsePublishedAt(item.PubDate)
        if err != nil {
            log.Printf("Skipping publication date for %q: %v", item.Title, err)
            // Save the post anyway, with published_at set to NULL.
        }

        now := time.Now()

        err = s.db.CreatePost(ctx, database.CreatePostParams{
            ID:        uuid.New(),
            CreatedAt: now,
            UpdatedAt: now,
            Title:     item.Title,
            Url:       item.Link,
            Description: sql.NullString{
                String: item.Description,
                Valid:  item.Description != "",
            },
            PublishedAt: publishedAt,
            FeedID:      fetchedFeed.ID,
        })
        if err != nil {
            log.Printf("Failed to save post %q (%s): %v", item.Title, item.Link, err)
            continue
        }
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
