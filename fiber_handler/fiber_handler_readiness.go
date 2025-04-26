package fiber_handler

import "github.com/gofiber/fiber/v2"

func FiberHandlerReadiness(f *fiber.Ctx) error {
	return FiberRespondWithJSON(f, 200, struct{}{})
}
