package validation

import (
	"boilerplate/exception"
	"boilerplate/model"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"strconv"
	"strings"
	"time"
)

func ValidateDate(value interface{}) error {
	_, err := time.Parse("2006-01-02", value.(string))
	if err != nil {
		return errors.New(model.NOT_VALID_ERR_TYPE)
	}
	return nil
}

func ValidateDateTime(value interface{}) error {
	_, err := time.Parse("2006-01-02 15:04:05", value.(string))
	if err != nil {
		return errors.New(model.NOT_VALID_ERR_TYPE)
	}
	return nil
}

func linkUrlValidation(linkType string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if linkType == "1" && s == "" {
			return exception.NOT_VALID
		}
		return nil
	}
}
func menuIdValidation(linkType string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if linkType == "0" && s == "" {
			return exception.NOT_VALID
		}
		return nil
	}
}
func cashbackTypeValidation() validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if s != "AMOUNT" && s != "PERCENTAGE" {
			return exception.NOT_VALID
		}
		return nil
	}
}

func voucherTypeValidation() validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if s != "GLOBAL" && s != "NEW-USER" && s != "OLD-USER" && s != "SPECIFIC-USER" {
			return exception.NOT_VALID
		}
		return nil
	}
}

func sortTypeValidation() validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		s = strings.ToLower(s)

		if s == "" {
			return nil
		}

		if s == "asc" || s == "desc" {
			return nil
		}

		return exception.NOT_VALID
	}
}

func priorityValidation() validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		a, err := strconv.Atoi(s)
		if err != nil {
			return exception.NOT_VALID
		}
		length := 0
		for a != 0 {
			if length >= 3 {
				return errors.New(model.Max(999))
			}
			a /= 10
			length = length + 1
		}

		return nil
	}
}

func npsSortFieldValidation() validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		s = strings.ToLower(s)

		if s == "" {
			return nil
		}

		switch s {
		case "agent_id":
			return nil
		case "score":
			return nil
		case "notes":
			return nil
		case "created_at":
			return nil
		case "updated_at":
			return nil
		default:
			return exception.NOT_VALID
		}
	}
}

func articleSortFieldValidation() validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		s = strings.ToLower(s)

		if s == "" {
			return nil
		}

		switch s {
		case "title":
			return nil
		case "label_name":
			return nil
		case "link_url":
			return nil
		case "created_at":
			return nil
		case "updated_at":
			return nil
		default:
			return exception.NOT_VALID
		}
	}
}
