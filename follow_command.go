package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Bis-sonido/gator/internal/database"
	"github.com/google/uuid"
)

func follow(s *state, cmd command, user database.User) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("must provide url")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	result, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}

	fmt.Printf("Feed: %s\n", result.FeedName)
	fmt.Printf("User: %s\n", result.UserName)
	return nil
}

func following(s *state, _ command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to fetch feed follows: %v", err)
	}

	if len(follows) == 0 {
		fmt.Println("You are not following any feeds")
		return nil
	}

	fmt.Printf("Feeds followed by %s:\n", user.Name)
	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}
	return nil
}
