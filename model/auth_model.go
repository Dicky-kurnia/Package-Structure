package model

import "time"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type LogoutResponse struct {
	AccessToken string `json:"access_token"`
}

type JwtPayload struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}

type TokenDetails struct {
	AccessToken string
	AtExpires   int64
	RtExpires   int64
}

type GetUserProfileResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type SetDataRedis struct {
	Key           string
	Data          interface{}
	Exp           time.Duration
	IsExternalDel bool
}
