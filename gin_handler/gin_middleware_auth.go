package gin_handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/neWbie-saby/rss-aggregator/internal/auth"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
)

type AuthHandler func(*gin.Context, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler AuthHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey, err := auth.GetAPIKey(c.Request.Header)
		if err != nil {
			GinRespondWithError(c, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(c.Request.Context(), apiKey)
		if err != nil {
			GinRespondWithError(c, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(c, user)
	}
}
