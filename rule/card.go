package rule

import (
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateIranianCardNumber Luan algorithm for credit/card numbers
func ValidateIranianCardNumber(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return true
	}
	// strip spaces and dashes
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, " ", "")
	if len(s) != 16 {
		return false
	}
	sum := 0
	for i := 0; i < 16; i++ {
		d, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}
		// Luhn: from left, double every second digit (depending on convention).
		// We'll apply the typical algorithm by index from right.
		pos := 15 - i
		if pos%2 == 1 {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}
	return sum%10 == 0
}
