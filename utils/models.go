package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func DatabaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
		UserID:    dbFeed.Userid,
	}
}

func DatabaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, DatabaseFeedToFeed(dbFeed))
	}

	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func DatabaseFeedFollowToFeedFollow(dbFeedFollows database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.CreatedAt,
		UpdatedAt: dbFeedFollows.UpdatedAt,
		UserID:    dbFeedFollows.UserID,
		FeedID:    dbFeedFollows.FeedID,
	}
}

func DatabaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedfollows := []FeedFollow{}

	for _, dbFeedFollow := range dbFeedFollows {
		feedfollows = append(feedfollows, DatabaseFeedFollowToFeedFollow(dbFeedFollow))
	}

	return feedfollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func DatabasePostToPost(dbPost database.Post) Post {
	var desc *string
	if dbPost.Description.Valid {
		desc = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: desc,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func DatabasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}

	for _, post := range dbPosts {
		posts = append(posts, DatabasePostToPost(post))
	}

	return posts
}
