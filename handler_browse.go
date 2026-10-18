package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ashraddhansh/gator/internal/database"
)


func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)

	if len(cmd.args) > 1 {
		return fmt.Errorf("usage: browse <limit>")
	}

	if len(cmd.args) == 1 {
		parsedLimit, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil || parsedLimit <= 0 {
			return fmt.Errorf("limit must be a positive integer")
		}
		limit = int32(parsedLimit)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: limit,
	})
	if err != nil {
		return fmt.Errorf("Fetching post for user: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("%s\n", post.Title)
		fmt.Printf("%s\n", post.Description.String)
		fmt.Printf("%s\n", post.Url)
		fmt.Println("----")
	}
	return nil
}
