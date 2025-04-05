package controller

import (
	"github.com/15011106/LetsWork/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	userUc usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) *UserController {
	return &UserController{userUc: u}
}

func (c *UserController) CreateUser(ctx fiber.Ctx) error {

	return ctx.SendString("implement me")
}
func (c *UserController) DeleteUser(ctx fiber.Ctx) error {

	return ctx.SendString("implement me")
}
