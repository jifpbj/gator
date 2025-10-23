package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jifpbj/gator/internal/database"
	"github.com/lib/pq"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("needs an argument for %s", cmd.Name)
	}

	ctx := context.Background()

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	user, err := s.db.CreateUser(ctx, userParams)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				fmt.Fprintf(os.Stderr, "User with name %s already exists\n", userParams.Name)
				os.Exit(1)
			}
		}
		return fmt.Errorf("couldn't create user: %w", err)
	}

	s.cfg.SetUser(user.Name)

	fmt.Printf("User %s created with ID %s at %s, updated at %s\n", user.Name, user.ID, user.CreatedAt, user.UpdatedAt)
	return nil
}
