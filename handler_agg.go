package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/IgorP25/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_requests>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Could not fetch next feeds", err)
		return
	}

	log.Println("Found a feed to fetch!")

	scrapeFeed(s.db, nextFeed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Could not mark feed %s as fetched: %v", feed.Name, err)
		return
	}

	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Could not fetch feed %s: %v", feed.Name, err)
		return
	}

	savePosts(db, feed, fetchedFeed)

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(fetchedFeed.Channel.Item))
}

func printFeed(feed *RSSFeed) {
	fmt.Printf("Channel: %s\n", feed.Channel.Title)
	// fmt.Printf("Link: %s\n", feed.Channel.Link)
	// fmt.Printf("Description: %s\n", feed.Channel.Description)
	fmt.Println("")

	for _, item := range feed.Channel.Item {
		fmt.Printf("\tTitle: %s\n", item.Title)
		// fmt.Printf("\tLink: %s\n", item.Link)
		// fmt.Printf("\tDescription: %s\n", item.Description)
		// fmt.Printf("\tPubDate: %s\n", item.PubDate)
		fmt.Println("")
	}
}

func savePosts(db *database.Queries, feedDB database.Feed, feed *RSSFeed) {
	for _, item := range feed.Channel.Item {
		publishedat := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedat = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		currentTime := time.Now()
		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedat,
			FeedID:      feedDB.ID,
		})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "unique_violation":
					continue
				}
			} else {
				log.Printf("Could not save post %s: %v", item.Title, err)
				continue
			}
		}
	}
}
