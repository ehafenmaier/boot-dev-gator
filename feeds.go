package main

import (
	"context"
	"fmt"
	"github.com/ehafenmaier/boot-dev-gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	// Get the current user from the database
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	// Create feed in the database
	dbParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), dbParams)
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	fmt.Printf("Feed Added\nName: %s\nUrl: %s\n", feed.Name, feed.Url)

	return nil
}

func handlerAllFeeds(s *state, _ command) error {
	// Get all feeds from the database
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}

	// Print all feeds
	for _, feed := range feeds {
		fmt.Printf("Name: %s\nUrl: %s\nUser: %s\n\n", feed.Name, feed.Url, feed.UserName)
	}

	return nil
}

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: follow <url>")
	}

	// Get the current user from the database
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	// Get the feed from the database
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}

	dbParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), dbParams)
	if err != nil {
		return fmt.Errorf("error following feed: %v", err)
	}

	fmt.Printf("Feed Followed\nFeed: %s\nUser: %s\n", feedFollow.FeedName, feedFollow.UserName)

	return nil
}
