package main

import (
	"context"
	"fmt"
)

func getUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %v", err)
	}

	userString := s.cfg.CurrentUserName

	for _, user := range users {
		if user.Name == userString {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
