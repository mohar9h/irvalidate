package rule

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var postalRegex = regexp.MustCompile(`^\d{10}$`)

func ValidateIranianPostalCode(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return true
	}
	return postalRegex.MatchString(s)
}
