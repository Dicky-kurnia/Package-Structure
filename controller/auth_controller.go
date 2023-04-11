package controller

import (
	"boilerplate/exception"
	"boilerplate/middleware"
	"boilerplate/model"
	"boilerplate/service"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type authController struct {
	service service.UserService
}

func NewAuthController(service service.UserService) *authController {
	return &authController{service}
}

func (controller *authController) Route(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Post("", controller.Login)
	auth.Post("/logout", middleware.CheckToken, controller.Logout)
}

func (controller *authController) Login(c *fiber.Ctx) error {
	request := new(model.LoginRequest)

	if err := c.BodyParser(request); err != nil {
		exception.PanicIfNeeded(err)
	}

	response, err := controller.service.Login(request)
	exception.PanicIfNeeded(err)

	return c.Status(200).JSON(model.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}

func (controller *authController) Logout(c *fiber.Ctx) error {
	tokenSlice := strings.Split(c.Get("Authorization"), "Bearer ")

	var tokenString string
	if len(tokenSlice) == 2 {
		tokenString = tokenSlice[1]
	}

	err := controller.service.Logout(tokenString)

	if err != nil {
		exception.PanicIfNeeded(err)
	}

	return c.Status(200).JSON(model.Response{
		Code:   200,
		Status: "OK",
	})
}
