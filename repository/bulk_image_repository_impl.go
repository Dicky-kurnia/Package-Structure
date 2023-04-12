package repository

import (
	"boilerplate/model"

	"gorm.io/gorm"
)

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) ImageRepository {
	return &imageRepository{db: db}
}

func (ir *imageRepository) Save(img *model.Image) error {
	return ir.db.Create(img).Error
}

func (ir *imageRepository) GetImageByName(name string) (*model.Image, error) {
	var image model.Image
	if err := ir.db.Where("name = ?", name).First(&image).Error; err != nil {
		return nil, err
	}
	return &image, nil
}
