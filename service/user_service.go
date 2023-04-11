package service

import "boilerplate/model"

type UserService interface {
	Login(request *model.LoginRequest) (response *model.LoginResponse, err error)
	Logout(token string) error
}
