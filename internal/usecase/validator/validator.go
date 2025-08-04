package validator

import (
	"bakery-api/internal/usecase/dto"
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

const nonSpecialCharRegex = `^[a-zA-Z0-9\s]+$`

func ValidateNonSpecialCharacter(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	matched, _ := regexp.MatchString(nonSpecialCharRegex, name)
	return matched
}

func GetValidateError(err error) []dto.APIError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]dto.APIError, len(ve))
		for i, fe := range ve {
			out[i] = dto.APIError{Code: fe.Field(), Message: msgForTag(fe.Tag())}
		}
		return out
	}
	return []dto.APIError{}
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "non_special_char":
		return "Can not containss special character"
	}
	return "invalid input"
}
