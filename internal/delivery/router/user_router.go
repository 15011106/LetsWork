package router

import (
	"github.com/15011106/LetsWork/internal/delivery/controller"
	"github.com/15011106/LetsWork/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

func RegisterUserRoutes(app *fiber.App, userUsecase usecase.UserUsecase) {

	uc := controller.NewUserController(userUsecase)

	userApi := app.Group("user")
	userApi.Post("/create-user", uc.CreateUser)
	userApi.Post("/delete-user", uc.DeleteUser)
}
