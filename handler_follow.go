package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jifpbj/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enough args")
	}

	url := cmd.Args[0]
	feedID, err := s.db.GetFeedIDFromURL(context.Background(), url)
	if err != nil {
		return err
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	followRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedID,
	})
	if err != nil {
		return fmt.Errorf("error creating follow feed: %w", err)
	}
	fmt.Println("Following feed:", followRow.FeedName)
	fmt.Println("Assigned to user:", followRow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, row := range feedFollows {
		fmt.Println(row.FeedName)
	}
	return nil
}
