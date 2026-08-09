package main

import (
	"fmt"
	"time"
)

func aggCommand(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <duration>", cmd.name)
	}

	time_between_reqs := cmd.args[0]

	duration, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("invalid duration: %w", duration)
	}
	ticker := time.NewTicker(duration)
	fmt.Printf("Collecting feeds every %v...", duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}
