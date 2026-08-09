package main

import (
	"context"
	"fmt"
)

func feeds(s *state, _ command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %v\n", feed.FeedName)
		fmt.Printf("* %v\n", feed.Url)
		fmt.Printf("* %v\n", feed.UserName)
	}

	return nil
}
