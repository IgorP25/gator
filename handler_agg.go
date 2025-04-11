package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	printFeed(feed)

	return nil
}

func printFeed(r *RSSFeed) {
	fmt.Printf("Channel: %s\n", r.Channel.Title)
	fmt.Printf("Link: %s\n", r.Channel.Link)
	fmt.Printf("Description: %s\n", r.Channel.Description)
	fmt.Println("")

	for _, item := range r.Channel.Item {
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Description: %s\n", item.Description)
		fmt.Printf("PubDate: %s\n", item.PubDate)
		fmt.Println("")
	}
}
