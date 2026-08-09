package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Bis-sonido/gator/internal/database"
	"github.com/google/uuid"
)

func registerUser(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("invalid command")
	}
	name := cmd.args[0]

	ctx := context.Background()

	now := time.Now()

	user, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	setUser := s.cfg.SetUser(name)
	if setUser != nil {
		return fmt.Errorf("failed to set user in config: %v", setUser)
	}

	fmt.Printf("User '%s' registered successfully.\n", user.Name)
	log.Printf("Debug - Created user: %+v", user)

	return nil
}
