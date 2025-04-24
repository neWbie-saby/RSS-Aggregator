package chi_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
	"github.com/neWbie-saby/rss-aggregator/utils"
)

func (apiCfg *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 200, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedfollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		RespondWithError(w, 201, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	RespondWithJSON(w, 200, utils.DatabaseFeedFollowToFeedFollow(feedfollow))
}

func (apiCfg *ApiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		RespondWithError(w, 201, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	RespondWithJSON(w, 200, utils.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *ApiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		RespondWithError(w, 201, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		RespondWithError(w, 201, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	RespondWithJSON(w, 200, struct{}{})
}
