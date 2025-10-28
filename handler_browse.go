package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jifpbj/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2

	if len(cmd.Args) == 1 {
		if result, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = result
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error retrieving posts:%w", err)
	}
	fmt.Println("Printing posts for user:", user.Name)
	printPosts(posts)
	return nil
}

func printPosts(posts []database.Post) {
	for _, post := range posts {
		fmt.Println("========Start Post========")
		fmt.Println("Title:", post.Title)
		fmt.Println("url:", post.Url)
		fmt.Println("published:", post.PublishedAt)
		fmt.Println("description", post.Description)
		fmt.Println("==========End Post========")
	}
}
