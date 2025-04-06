package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {

		if err := app.Listen(":8080"); err != nil {
			log.Panicf("Server failed to start: %v", err)
		}
	}()

	<-c
	log.Println("Gracefully shutting down...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	if err := db.Close(); err != nil {
		log.Printf("Error closing DB: %v", err)
	}

	log.Println("Server gracefully stopped")

	log.Fatal(app.Listen(":8080"))
}
