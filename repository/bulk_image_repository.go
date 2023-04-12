package repository

import "boilerplate/model"

type ImageRepository interface {
	Save(img *model.Image) error
	GetImageByName(name string) (*model.Image, error)
}
