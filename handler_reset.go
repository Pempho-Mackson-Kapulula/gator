package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not delete users: %w", err)
	}

	fmt.Println("Successfully deleted all users")

	return nil
}
