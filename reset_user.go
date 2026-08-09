package main

import (
	"context"
	"fmt"
)

func resetUser(s *state, _ command) error {

	err := s.db.DeleteUser(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	fmt.Printf("user reset successfully.\n")
	return nil
}
