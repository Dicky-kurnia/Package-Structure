package validation

import (
	"boilerplate/exception"
	"boilerplate/model"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/goccy/go-json"
)

func ValidateLogin(request model.LoginRequest) error {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username,
			validation.Required.Error(model.NOT_BLANK_ERR_TYPE),
		),
		validation.Field(&request.Password,
			validation.Required.Error(model.NOT_BLANK_ERR_TYPE),
		),
	)

	if err != nil {
		b, _ := json.Marshal(err)
		err = exception.ValidationError{
			Message: string(b),
		}
		return err
	}

	return nil
}
