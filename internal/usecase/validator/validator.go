package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const nonSpecialCharRegex = `^[a-zA-Z0-9\s]+$`

func ValidateNonSpecialCharacter(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	matched, _ := regexp.MatchString(nonSpecialCharRegex, name)
	return matched
}

func GetValidateError(err *validator.ValidationErrors) map[string]string {
	out := make(map[string]string)
	for _, fe := range *err {
		out[fe.Field()] = msgForTag(fe.Tag(), fe.Param())
	}
	return out
}

func msgForTag(tag string, param string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "non_special_char":
		return "Can not contains special character"
	case "max":
		return "Can not have more than " + param + " characters"
	}
	return "invalid input"
}
