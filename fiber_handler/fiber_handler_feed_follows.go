package fiber_handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
	"github.com/neWbie-saby/rss-aggregator/utils"
)

func (apiCfg *ApiConfig) HandlerCreateFeedFollow(f *fiber.Ctx, user database.User) error {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := parameters{}
	if err := f.BodyParser(&params); err != nil {
		return FiberRespondWithError(f, 200, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	feedfollow, err := apiCfg.DB.CreateFeedFollow(f.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		return FiberRespondWithError(f, 201, fmt.Sprintf("Couldn't create feed follow: %v", err))
	}

	return FiberRespondWithJSON(f, 200, utils.DatabaseFeedFollowToFeedFollow(feedfollow))
}

func (apiCfg *ApiConfig) HandlerGetFeedFollows(f *fiber.Ctx, user database.User) error {
	feedFollows, err := apiCfg.DB.GetFeedFollows(f.Context(), user.ID)
	if err != nil {
		return FiberRespondWithError(f, 201, fmt.Sprintf("Couldn't get feed follows: %v", err))
	}

	return FiberRespondWithJSON(f, 200, utils.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *ApiConfig) HandlerDeleteFeedFollow(f *fiber.Ctx, user database.User) error {
	feedFollowIDStr := f.Params("feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		return FiberRespondWithError(f, 201, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
	}

	err = apiCfg.DB.DeleteFeedFollow(f.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		return FiberRespondWithError(f, 201, fmt.Sprintf("Couldn't delete feed follow: %v", err))
	}
	return FiberRespondWithJSON(f, 200, struct{}{})
}
