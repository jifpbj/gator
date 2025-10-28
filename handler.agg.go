package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jifpbj/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("agg needs 1 arg; time between requests")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error parsing time between requests: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	if err := scrapeFeeds(s); err != nil {
		return err
	}
	for {
		<-ticker.C
		if err := scrapeFeeds(s); err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}

	// get RSS feed
	RSSFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error getting RSS from next Feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}

	fmt.Println("Printing from Feed:", nextFeed.Name)
	for _, item := range RSSFeed.Channel.Item {
		timePub, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return fmt.Errorf("error parsing time: %w", err)
		}

		fmt.Println("Saving:  ", item.Title)
		fmt.Println("time:", item.PubDate)
		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       sql.NullString{String: item.Title, Valid: true},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: timePub,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			return fmt.Errorf("error creating post: %w", err)
		}
		fmt.Println("Post created Successfully")
		fmt.Printf("%+v\n", post)

	}
	fmt.Println("============End of Feed============")
	fmt.Println("===================================")
	return nil
}
