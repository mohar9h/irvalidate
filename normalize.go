package irvalidate

import "strings"

// Persian (Eastern Arabic) digits -> ASCII digits map
var persianDigitMap = map[rune]rune{
	'۰': '0',
	'۱': '1',
	'۲': '2',
	'۳': '3',
	'۴': '4',
	'۵': '5',
	'۶': '6',
	'۷': '7',
	'۸': '8',
	'۹': '9',
}

// NormalizePersianDigits returns a string where Persian digits are mapped to
// ASCII digits. It leaves other characters untouched.
func NormalizePersianDigits(input string) string {
	if string(input) == "" {
		return input
	}
	output := strings.Builder{}

	output.Grow(len(input))

	for _, r := range input {
		if v, ok := persianDigitMap[r]; ok {
			output.WriteRune(v)
		} else {
			output.WriteRune(r)
		}
	}
	return output.String()
}

// TrimSpaces normalizes whitespace (simple helper)
func TrimSpaces(input string) string {
	return strings.TrimSpace(input)
}
