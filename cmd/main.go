package main

import (
	"log"

	"github.com/15011106/LetsWork/internal/delivery/router"
	"github.com/15011106/LetsWork/internal/infrastructure"
	repository_impl "github.com/15011106/LetsWork/internal/repository/repository_impl/user"
	"github.com/15011106/LetsWork/internal/usecase/usecase_impl"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	db, err := infrastructure.NewDatabase()
	if err != nil {
		panic(err)
	}

	userRepository := repository_impl.NewUserRepository(db.Rds)
	userUsecase := usecase_impl.NewUserUsecase(userRepository)

	router.RegisterSystemRoutes(app)
	router.RegisterUserRoutes(app, userUsecase)

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
