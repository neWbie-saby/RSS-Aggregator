package fiber_handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/neWbie-saby/rss-aggregator/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func FiberRespondWithJSON(f *fiber.Ctx, code int, payload interface{}) error {
	f.Status(code)
	return f.JSON(payload)
}

func FiberRespondWithError(f *fiber.Ctx, code int, msg string) error {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	return FiberRespondWithJSON(f, code, errResponse{
		Error: msg,
	})
}
