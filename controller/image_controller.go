package controller

import (
	"github.com/gofiber/fiber/v2"
)

type ImageController interface {
	GetImageByName(c *fiber.Ctx) error
	Route(group fiber.Router)
}
