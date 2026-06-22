package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Pempho-Mackson-Kapulula/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username can not be empty")
	}

	//handle duplicate users
	_, err := s.queries.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		fmt.Println("user does not exist")
		os.Exit(1)
	}

	// set user
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("could not set user: %v", err)
	}
	fmt.Printf("User has been set\n")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username cannot be empty")
	}

	//fill in user details
	user, err := s.queries.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})

	if err != nil {
		fmt.Println("user already exists:", err)
		os.Exit(1)
	}

	// set current user in config
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}
	// print success message
	fmt.Println("User registered successfully!")

	return nil
}
