package main

import (
	"fmt"
	"context"
	"time"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/Bis-sonido/gator/internal/database"
)

func scrapeFeeds(s *state) error {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get next feed: %v", err)
	}

	_, err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark the feed: %v", err)
	}

	fetchFeedRss, fError := fetchFeed(context.Background(), nextFeed.Url)
	if fError != nil {
		return fmt.Errorf("failed to fetch the feed: %v", fError)
	}

	
	for _, feed := range fetchFeedRss.Channel.Item{
		parsedPublishedAt, _ := ParsePublishedAt(feed.PubDate)
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: feed.Title,
			Url: feed.Link,
			Description: sql.NullString{String: feed.Description, Valid: true},
			PublishedAt: parsedPublishedAt,
			FeedID: nextFeed.ID,
		})
		if err != nil {
			var pqErr *pq.Error
			if ok := errors.As(err, &pqErr); ok && pqErr.Code == "23505" {
				fmt.Printf("Post already exists: %v\n", feed.Title)
			} else {
				fmt.Printf("Failed to create post: %v\n", err)
			}
			continue
		}
	}
	return nil
}

func ParsePublishedAt(pubDate string) (sql.NullTime, error) {
	if pubDate == "" {
		return sql.NullTime{Valid: false}, nil
	}

	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.RFC3339,
	}

	for _, layout := range layouts {
		if parsedTime, err := time.Parse(layout, pubDate); err == nil {
			return sql.NullTime{Time: parsedTime, Valid: true}, nil
		}
	}

	return sql.NullTime{Valid: false}, nil
}