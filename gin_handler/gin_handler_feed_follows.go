package gin_handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
	"github.com/neWbie-saby/rss-aggregator/utils"
)

func (apiCfg *ApiConfig) HandlerCreateFeedFollow(c *gin.Context, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(c.Request.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		GinRespondWithError(c, 200, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedfollow, err := apiCfg.DB.CreateFeedFollow(c.Request.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		GinRespondWithError(c, 201, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	GinRespondWithJSON(c, 200, utils.DatabaseFeedFollowToFeedFollow(feedfollow))
}

func (apiCfg *ApiConfig) HandlerGetFeedFollows(c *gin.Context, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(c.Request.Context(), user.ID)
	if err != nil {
		GinRespondWithError(c, 201, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	GinRespondWithJSON(c, 200, utils.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *ApiConfig) HandlerDeleteFeedFollow(c *gin.Context, user database.User) {
	feedFollowIDStr := c.Param("feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		GinRespondWithError(c, 201, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(c.Request.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		GinRespondWithError(c, 201, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	GinRespondWithJSON(c, 200, struct{}{})
}
