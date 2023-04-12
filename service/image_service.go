package service

import (
	"boilerplate/model"
	"mime/multipart"
)

type ImageService interface {
	UploadImage(file multipart.File, header *multipart.FileHeader) error
	GetImageByName(name string) (*model.Image, error)
}
