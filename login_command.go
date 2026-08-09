package main

import (
	"context"
	"fmt"
	"github.com/Bis-sonido/gator/internal/config"
	"github.com/Bis-sonido/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("username is required")
	}

	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("failed to get user: %v", err)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}

	fmt.Printf("User '%s' logged in successfully.\n", username)
	return nil
}
