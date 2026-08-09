package main

import (
	"context"
	"fmt"
	"strconv"
	"github.com/Bis-sonido/gator/internal/database"
)

func browse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		var err error
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %v", err)
		}
		if limit <= 0 {
			return fmt.Errorf("limit must be greater than 0")
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to get posts: %v", err)
	}

	fmt.Printf("Found %d posts for user %s: \n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("URL: %s\n", post.Url)
	}
	return nil
}