package main

import (
	"context"
	"fmt"
	"time"

	"github.com/IgorP25/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]
	currentTime := time.Now()

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	fmt.Println("Added feed: ")
	fmt.Printf("\tChannel: %s\n", feed.Name)
	fmt.Printf("\tLink: %s\n", feed.Url)

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not register follow: %w", err)
	}

	fmt.Printf("%s is now following %s\n", user.Name, feed.Name)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Println("Saved feeds: ")
	for _, feed := range feeds {
		fmt.Printf("\tChannel: %s\n", feed.Name)
		fmt.Printf("\tLink: %s\n", feed.Url)
		fmt.Printf("\tUser: %s\n", feed.UserName)
	}

	return nil
}
