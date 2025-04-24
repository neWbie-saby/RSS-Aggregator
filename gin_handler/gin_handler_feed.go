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

func (apiCfg *ApiConfig) HandlerCreateFeed(c *gin.Context, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(c.Request.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		GinRespondWithError(c, 200, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(c.Request.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		Userid:    user.ID,
	})
	if err != nil {
		GinRespondWithError(c, 201, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	GinRespondWithJSON(c, 200, utils.DatabaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) HandlerGetFeeds(c *gin.Context) {
	feeds, err := apiCfg.DB.GetFeeds(c.Request.Context())
	if err != nil {
		GinRespondWithError(c, 201, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	GinRespondWithJSON(c, 200, utils.DatabaseFeedsToFeeds(feeds))
}
