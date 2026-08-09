package main

import (
	"context"
	"fmt"

	"github.com/Bis-sonido/gator/internal/database"
)

func unfollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("must provide a feed url")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete feed: %v", err)
	}

	fmt.Printf("succesfully unfollowed feed: %s\n", feed.Name)
	return nil
}
