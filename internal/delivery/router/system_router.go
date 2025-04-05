package router

import "github.com/gofiber/fiber/v3"

type errorDto struct {
	Err       string `json:"error"`
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
}

func RegisterSystemRoutes(app *fiber.App) {
	app.Get("/ping", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Use(func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(errorDto{
			Err:       "Not Found",
			Message:   "No route matched",
			ErrorCode: fiber.StatusNotFound,
		})
	})
}
