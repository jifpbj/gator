package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jifpbj/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enough args")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	followRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating follow feed: %w", err)
	}

	fmt.Println("Following feed:", followRow.FeedName)
	fmt.Println("Assigned to user:", followRow.UserName)
	return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user")
		return nil
	}

	for _, row := range feedFollows {
		fmt.Println("*", row.FeedName)
	}
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
