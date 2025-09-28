package rule

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Accepts Persian/Arabic script characters, Persian zero-width non-joiner (\u200c) and spaces.
var persianTextRegex = regexp.MustCompile(`^[\p{Arabic}\s\u200c]+$`)

func ValidatePersianText(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return true
	}
	return persianTextRegex.MatchString(s)
}
