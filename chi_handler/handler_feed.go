package chi_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
	"github.com/neWbie-saby/rss-aggregator/utils"
)

func (apiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 200, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		Userid:    user.ID,
	})
	if err != nil {
		RespondWithError(w, 201, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	RespondWithJSON(w, 200, utils.DatabaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		RespondWithError(w, 201, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	RespondWithJSON(w, 200, utils.DatabaseFeedsToFeeds(feeds))
}
