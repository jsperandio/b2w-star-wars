package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberMiddleware struct{}

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func InitFiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// Add Cache layer
		// cache.New(),
		// Add simple logger.
		logger.New(),
	)
}
