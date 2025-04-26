package fiber_handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/neWbie-saby/rss-aggregator/internal/auth"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
)

type AuthHandler func(*fiber.Ctx, database.User) error

func (cfg *ApiConfig) MiddlewareAuth(handler AuthHandler) fiber.Handler {
	return func(f *fiber.Ctx) error {
		apiKey, err := auth.GetAPIKey(f.GetReqHeaders())
		if err != nil {
			return FiberRespondWithError(f, 403, fmt.Sprintf("Auth error: %v", err))
		}

		user, err := cfg.DB.GetUserByAPIKey(f.Context(), apiKey)
		if err != nil {
			return FiberRespondWithError(f, 400, fmt.Sprintf("Couldn't get user: %v", err))
		}

		return handler(f, user)
	}
}
