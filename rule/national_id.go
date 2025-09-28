package rule

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

// ValidateIranianNationalCode validates Iranian national codes (کد ملی).
// Algorithm: 10 digits, and the 10th digit is the checksum of the first 9 digits.
func ValidateIranianNationalCode(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return true
	}
	if len(s) != 10 {
		return false
	}
	// prevent obvious invalid sequences like all same digit
	allSame := true
	for i := 1; i < 10; i++ {
		if s[i] != s[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		n, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}
		sum += n * (10 - i)
	}
	check, err := strconv.Atoi(string(s[9]))
	if err != nil {
		return false
	}
	rem := sum % 11
	if rem < 2 {
		return check == rem
	}
	return check == 11-rem
}
