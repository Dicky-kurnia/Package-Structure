package service

import (
	"boilerplate/exception"
	"boilerplate/helper"
	"boilerplate/model"
	"boilerplate/repository"
	"boilerplate/validation"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

type userService struct {
	UserRepository repository.UserRepository
}

func (service userService) Login(request *model.LoginRequest) (response *model.LoginResponse, err error) {
	if err = validation.ValidateLogin(*request); err != nil {
		return response, err
	}

	user, err := service.UserRepository.GetUser(request.Username)
	if user.Username == "" || err != nil {
		return nil, exception.USERNAME_OR_PASSWORD_INVALID
	}

	if os.Getenv("IS_PRODUCTION") == "1" {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			return nil, exception.USERNAME_OR_PASSWORD_INVALID
		}
	} else {
		if strings.Compare(user.Password, request.Password) != 0 {
			return nil, exception.USERNAME_OR_PASSWORD_INVALID
		}
	}

	jwtPayload := model.JwtPayload{
		UserId:   user.Id.Hex(),
		Username: user.Username,
	}

	//generate token
	ts := helper.CreateToken(jwtPayload)

	//save metadata to redis
	helper.CreateAuth(jwtPayload, ts)

	response = &model.LoginResponse{
		AccessToken: ts.AccessToken,
	}

	return response, nil
}

func (service userService) Logout(token string) error {
	helper.DelRedis(token)

	return nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{UserRepository: userRepository}
}
