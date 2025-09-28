package rule

import (
	"regexp"

	"github.com/mohar9h/irvalidate"
)

var persianAlphaRE = regexp.MustCompile(`^[\u0600-\u06FF\u0750-\u077F\uFB50-\uFDFF\uFE70-\uFEFF\u200C\s]+$`)

// IsPersianAlpha returns true if the string contains only Persian/Arabic script
// letters and optional ZWNJ and whitespace.
func IsPersianAlpha(s string) bool {
	s = irvalidate.TrimSpaces(s)
	if s == "" {
		return false
	}
	return persianAlphaRE.MatchString(s)
}
