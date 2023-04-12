package model

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Name string
	Size int64
	URL  string
}
