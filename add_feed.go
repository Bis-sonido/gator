package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Bis-sonido/gator/internal/database"
	"github.com/google/uuid"
)

func addFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("invalid command")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	now := time.Now().UTC()

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create record: %v", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to follow created feed: %w", err)
	}

	fmt.Printf("feed: %v", feed)

	return nil

}
