package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Pempho-Mackson-Kapulula/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("no enough arguments for this command")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	//fill in feed details
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
    	return fmt.Errorf("could not create feed follow: %w", err)
	}

	fmt.Println("Feed created successfully!")
	fmt.Printf("ID: %s\n", feed.ID)
	fmt.Printf("Created at: %s\n", feed.CreatedAt)
	fmt.Printf("Udated at: %s\n", feed.UpdatedAt)
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("Url: %s\n", feed.Url)
	fmt.Printf("User ID: %s\n", feed.UserID)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	//get feeds
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	for _, f := range feeds {
		fmt.Printf("Username: %s\n", f.UserName)
		fmt.Printf("URL: %s\n", f.Url)
		fmt.Printf("Feed Name: %s\n", f.FeedName)
	}

	return nil
}
