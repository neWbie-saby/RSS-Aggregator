package fiber_handler

import "github.com/gofiber/fiber/v2"

func FiberHandlerErr(f *fiber.Ctx) error {
	return FiberRespondWithError(f, 400, "Something went wrong")
}
