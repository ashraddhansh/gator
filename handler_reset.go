package main

import (
	"context"
	"fmt"
)


func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error deleting the users: %v", err)
	}
	fmt.Println("All the users have been deleted!")
	return nil
}
