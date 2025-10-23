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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		os.Exit(1)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("Username set to %s\n", s.cfg.CurrentUserName)
	return nil
}
