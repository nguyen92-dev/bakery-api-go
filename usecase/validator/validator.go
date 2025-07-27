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
