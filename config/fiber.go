package config

import (
	"boilerplate/exception"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: exception.ErrorHandler,
	}
}
