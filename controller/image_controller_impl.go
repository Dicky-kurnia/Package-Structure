package controller

import (
	"boilerplate/model"
	"boilerplate/service"

	"github.com/gofiber/fiber/v2"
)

type imageController struct {
	service service.ImageService
}

func NewImageController(service service.ImageService) ImageController {
	return &imageController{service: service}
}

func (controller *imageController) Route(group fiber.Router) {
	group.Post("", controller.UploadImage)
	group.Get(":image_name", controller.GetImageByName)
}

func (controller *imageController) UploadImage(c *fiber.Ctx) error {
	// Get the uploaded file
	fileHeader, err := c.FormFile("bulk_file")
	if err != nil {
		return err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	// Upload the image
	err = controller.service.UploadImage(file, fileHeader)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(model.Response{
		Code:   200,
		Status: "OK",
	})
}

func (controller *imageController) GetImageByName(c *fiber.Ctx) error {
	name := c.Params("image_name")
	image, err := controller.service.GetImageByName(name)
	if err != nil {
		return err
	}

	return c.JSON(image)
}
