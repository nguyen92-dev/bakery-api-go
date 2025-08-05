package validator

import (
	"bakery-api/internal/usecase/dto"
	"regexp"

	"github.com/go-playground/validator/v10"
)

const nonSpecialCharRegex = `^[a-zA-Z0-9\s]+$`

func ValidateNonSpecialCharacter(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	matched, _ := regexp.MatchString(nonSpecialCharRegex, name)
	return matched
}

func GetValidateError(err *validator.ValidationErrors) []dto.APIError {
	out := make([]dto.APIError, len(*err))
	for i, fe := range *err {
		out[i] = dto.APIError{Code: fe.Field(), Message: msgForTag(fe.Tag())}
	}
	return out
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "non_special_char":
		return "Can not contains special character"
	}
	return "invalid input"
}
