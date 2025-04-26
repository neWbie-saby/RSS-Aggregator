package fiber_handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
	"github.com/neWbie-saby/rss-aggregator/utils"
)

func (apiCfg *ApiConfig) HandlerCreateFeed(f *fiber.Ctx, user database.User) error {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	params := parameters{}
	if err := f.BodyParser(&params); err != nil {
		return FiberRespondWithError(f, 200, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	feed, err := apiCfg.DB.CreateFeed(f.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		Userid:    user.ID,
	})
	if err != nil {
		return FiberRespondWithError(f, 201, fmt.Sprintf("Couldn't create feed: %v", err))
	}

	return FiberRespondWithJSON(f, 200, utils.DatabaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) HandlerGetFeeds(f *fiber.Ctx) error {
	feeds, err := apiCfg.DB.GetFeeds(f.Context())
	if err != nil {
		return FiberRespondWithError(f, 201, fmt.Sprintf("Couldn't get feeds: %v", err))
	}

	return FiberRespondWithJSON(f, 200, utils.DatabaseFeedsToFeeds(feeds))
}
