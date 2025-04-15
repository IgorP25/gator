package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/IgorP25/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if !(len(cmd.Args) == 0 || len(cmd.Args) == 1) {
		return fmt.Errorf("usage: %s [limit]", cmd.Name)
	}

	limitString := "2"
	if len(cmd.Args) == 1 {
		limitString = cmd.Args[0]
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		return fmt.Errorf("could not parse limit: %w", err)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("could not get posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("Published: %v\n", post.PublishedAt.Time)
		fmt.Printf("From feed: %s\n", post.FeedName)
	}

	return nil
}
