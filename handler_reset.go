package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllPosts(context.Background())
	if err != nil {
		return fmt.Errorf("could not delete posts: %w", err)
	}

	err = s.db.DeleteAllFeedFollows(context.Background())
	if err != nil {
		return fmt.Errorf("could not delete feed follows: %w", err)
	}

	err = s.db.DeleteAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not delete feeds: %w", err)
	}

	err = s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not delete users: %w", err)
	}

	fmt.Println("Successfully deleted all data")

	return nil
}
