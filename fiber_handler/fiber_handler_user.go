package fiber_handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
	"github.com/neWbie-saby/rss-aggregator/utils"
)

func (apiCfg *ApiConfig) HandlerCreateUser(f *fiber.Ctx) error {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	if err := f.BodyParser(&params); err != nil {
		return FiberRespondWithError(f, 200, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	user, err := apiCfg.DB.CreateUser(f.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		return FiberRespondWithError(f, 201, fmt.Sprintf("Couldn't create user: %v", err))
	}

	return FiberRespondWithJSON(f, 200, utils.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandleGetUser(f *fiber.Ctx, user database.User) error {
	return FiberRespondWithJSON(f, 200, utils.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandleGetPostsForUser(f *fiber.Ctx, user database.User) error {
	posts, err := apiCfg.DB.GetPostsForUser(f.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		return FiberRespondWithError(f, 201, fmt.Sprintf("Couldn't get posts for user: %v", err))
	}
	return FiberRespondWithJSON(f, 200, utils.DatabasePostsToPosts(posts))
}
