package rule

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Mobile regex: starts with 09 and 10 digits total.
var mobileRegex = regexp.MustCompile(`^09\d{9}$`)

func ValidateIranianMobile(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return true // allow empty — let `required` tag handle presence
	}
	return mobileRegex.MatchString(s)
}
