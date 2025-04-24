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

func (apiCfg *ApiConfig) HandlerCreateUser(c *gin.Context) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(c.Request.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		GinRespondWithError(c, 200, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(c.Request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		GinRespondWithError(c, 201, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	GinRespondWithJSON(c, 200, utils.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandleGetUser(c *gin.Context, user database.User) {
	GinRespondWithJSON(c, 200, utils.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandleGetPostsForUser(c *gin.Context, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(c.Request.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		GinRespondWithError(c, 201, fmt.Sprintf("Couldn't get posts for user: %v", err))
		return
	}
	GinRespondWithJSON(c, 200, utils.DatabasePostsToPosts(posts))
}
