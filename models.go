package main

import (
	"time"

	"github.com/elBanna00/rss-agg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID	`json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string	`json:"name"`
	APIKey string  `json:"api_key"`
}

func databaseUsertoUser(dbUser database.User) User{
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		APIKey : dbUser.ApiKey	,
	}
}
type Feed struct {

	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func databaseFeedsToFeeds(DBFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _,dbfeed := range DBFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbfeed))
	}
	return feeds  
}

