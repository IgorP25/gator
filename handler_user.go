package main

import (
	"context"
	"fmt"
	"time"

	"github.com/IgorP25/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userName := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("user %s not found: %w", userName, err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User has been set to: %s\n", user.Name)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	cmdName := cmd.Args[0]

	currentTime := time.Now()

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmdName,
	})
	if err != nil {
		return fmt.Errorf("couldn't add user %s: %w", cmdName, err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %s has been registered.\n", user.Name)

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not reset users: %w", err)
	}

	fmt.Println("Users successfully reset.")

	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not get users: %w", err)
	}

	for _, user := range users {
		isAdmin := ""
		if user.Name == s.cfg.CurrentUserName {
			isAdmin = " (current)"
		}
		fmt.Printf("* %s%s\n", user.Name, isAdmin)
	}

	return nil
}
