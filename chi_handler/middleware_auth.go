package chi_handler

import (
	"fmt"
	"net/http"

	"github.com/neWbie-saby/rss-aggregator/internal/auth"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			RespondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
