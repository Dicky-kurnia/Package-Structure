package exception

import (
	"boilerplate/model"
	"errors"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type CustomErr struct {
	status  int
	code    string
	message string
}

func (c *CustomErr) Status() int {
	return c.status
}

func (c *CustomErr) Code() string {
	return c.code
}

func (c *CustomErr) Error() string {
	return c.message
}

func NewError(status int, code, message string) error {
	return &CustomErr{
		status:  status,
		code:    code,
		message: message,
	}
}

var (
	SLIDER_NOT_FOUND                = NewError(http.StatusBadRequest, model.BAD_REQUEST, "SLIDER_NOT_FOUND")
	PROMO_BANNER_NOT_FOUND          = NewError(http.StatusBadRequest, model.BAD_REQUEST, "PROMO_BANNER_NOT_FOUND")
	ARTICLE_NOT_FOUND               = NewError(http.StatusBadRequest, model.BAD_REQUEST, "ARTICLE_NOT_FOUND")
	DYNAMIC_CARD_NOT_FOUND          = NewError(http.StatusBadRequest, model.BAD_REQUEST, "DYNAMIC_CARD_NOT_FOUND")
	POPUPS_NOT_FOUND                = NewError(http.StatusBadRequest, model.BAD_REQUEST, "POPUPS_NOT_FOUND")
	CASHBACK_VOUCHER_NOT_FOUND      = NewError(http.StatusBadRequest, model.BAD_REQUEST, "CASHBACK_VOUCHER_NOT_FOUND")
	MENU_ID_NOT_FOUND               = NewError(http.StatusBadRequest, model.BAD_REQUEST, "MENU_ID_NOT_FOUND")
	USERNAME_OR_PASSWORD_INVALID    = NewError(http.StatusBadRequest, model.BAD_REQUEST, "USERNAME_OR_PASSWORD_INVALID")
	EXTENSION_NOT_ALLOWED           = NewError(http.StatusBadRequest, model.BAD_REQUEST, "EXTENSION_NOT_ALLOWED")
	NOT_VALID                       = NewError(http.StatusBadRequest, model.BAD_REQUEST, "NOT_VALID")
	START_AND_END_DATE_NOT_VALID    = NewError(http.StatusBadRequest, model.BAD_REQUEST, "START_AND_END_DATE_NOT_VALID")
	INVALID_EXCEL_FILE              = NewError(http.StatusBadRequest, model.BAD_REQUEST, "INVALID_EXCEL_FILE")
	STARTED_AT_EXPIRED_AT_NOT_VALID = NewError(http.StatusBadRequest, model.BAD_REQUEST, "STARTED_AT_EXPIRED_AT_NOT_VALID")
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var customError *CustomErr
	if errors.As(err, &customError) {
		return ctx.Status(customError.Status()).JSON(model.Response{
			Code:   customError.Status(),
			Status: customError.Code(),
			Error: map[string]interface{}{
				"general": customError.Error(),
			},
		})
	}
	_, ok := err.(ValidationError)
	if ok {
		var obj interface{}
		_ = json.Unmarshal([]byte(err.Error()), &obj)
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: model.BAD_REQUEST,
			Data:   nil,
			Error:  obj,
		})
	}
	if err.Error() == model.AUTHENTICATION_FAILURE_ERR_TYPE {
		return ctx.Status(401).JSON(model.Response{
			Code:   401,
			Status: model.UNAUTHORIZATION,
			Data:   nil,
			Error: map[string]interface{}{
				"general": model.AUTHENTICATION_FAILURE_ERR_TYPE,
			},
		})
	}

	return ctx.Status(500).JSON(model.Response{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
