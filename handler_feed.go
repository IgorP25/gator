package main

import (
	"context"
	"fmt"
	"time"

	"github.com/IgorP25/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	currentTime := time.Now()

	r, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("cannot create feed: %w", err)
	}

	fmt.Println(r)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	r, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("cannot get feeds: %w", err)
	}

	if len(r) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Println("")

	for _, feed := range r {
		fmt.Printf("Channel: %s\n", feed.Name)
		fmt.Printf("Link: %s\n", feed.Url)
		fmt.Printf("User: %s\n", feed.UserName)
	}

	return nil
}
