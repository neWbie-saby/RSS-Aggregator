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

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 200, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, 201, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	RespondWithJSON(w, 200, utils.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	RespondWithJSON(w, 200, utils.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandleGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		RespondWithError(w, 201, fmt.Sprintf("Couldn't get posts for user: %v", err))
		return
	}
	RespondWithJSON(w, 200, utils.DatabasePostsToPosts(posts))
}
