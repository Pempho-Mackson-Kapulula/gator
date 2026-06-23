package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Pempho-Mackson-Kapulula/gator/internal/database"
	"github.com/google/uuid"
)


func handlerFeedFollows (s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("please enter the feed URL only")
	}


	//get feed using url
	url := cmd.args[0]
	
	feedFromURL, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not get feed : %w", err)
	}

	//create feed and fill in feed details
	feed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID: 	feedFromURL.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create a feed follow : %w", err)
	}

	fmt.Println("Feed created successfully!")
	fmt.Printf("Username : %s\n", feed.UserName)
	fmt.Printf("Feed Name: %s\n", feed.FeedName)
	

	return nil
}

func handlerListFeedFollows (s *state, cmd command, user database.User) error{
	//get current user name
	
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil{
		return fmt.Errorf("could not get feed follows for user: %w", err)
	}

	for _,feedFollow := range feedFollows {
		fmt.Printf("- %s", feedFollow.FeedName)
	}

	return nil
}


func handlerUnfollow(s *state, cmd command, user database.User) error{
	if len(cmd.args) != 1 {
		return fmt.Errorf("please enter the feed URL only")
	}

	//get feed using url
	url := cmd.args[0]
	
	feedFromURL, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not get feed : %w", err)
	}

	err = s.db.DeleteFeedFollowForUser(context.Background(), database.DeleteFeedFollowForUserParams{
		UserID: user.ID,
		FeedID: feedFromURL.ID,
	})

	if err != nil {
		return fmt.Errorf("could not delete feed follow : %w", err)
	}

	return nil

}