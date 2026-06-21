package main

import (
	"fmt"
	"context"
)


func handlerReset(s *state, cmd command) error {
	err := s.queries.DeleteAllUsers(context.Background())
	if err != nil{
		return fmt.Errorf("Error: %v", err)
	}

	fmt.Println("Successfully deleted all users")

	return nil
}