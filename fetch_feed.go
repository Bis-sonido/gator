package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", "gator")
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch feed: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var dataFeed RSSFeed
	if err := xml.Unmarshal(data, &dataFeed); err != nil {
		return nil, fmt.Errorf("failed to unmarshal feed: %v", err)
	}

	dataFeed.Channel.Title = html.UnescapeString(dataFeed.Channel.Title)
	dataFeed.Channel.Description = html.UnescapeString(dataFeed.Channel.Description)

	for i := range dataFeed.Channel.Item {
		dataFeed.Channel.Item[i].Title = html.UnescapeString(dataFeed.Channel.Item[i].Title)
		dataFeed.Channel.Item[i].Description = html.UnescapeString(dataFeed.Channel.Item[i].Description)
	}

	return &dataFeed, nil

}
