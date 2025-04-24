package gin_handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func GinRespondWithJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}

func GinRespondWithError(c *gin.Context, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	GinRespondWithJSON(c, code, errResponse{
		Error: msg,
	})
}
