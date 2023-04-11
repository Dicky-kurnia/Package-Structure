package repository

import "boilerplate/entity"

type UserRepository interface {
	GetUser(username string) (entity.UserEntity, error)
}
