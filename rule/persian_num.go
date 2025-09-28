package rule

import (
	"regexp"

	"github.com/mohar9h/irvalidate"
)

// Accept strings of Persian digits (۰۱۲۳۴۵۶۷۸۹) OR ASCII digits depending on
// option. We precompile regexes for common patterns.
var (
	persianDigitsRE = regexp.MustCompile(`^[\x{06F0}-\x{06F9}]+$`)
	asciiDigitsRE   = regexp.MustCompile(`^[0-9]+$`)
	mixedDigitsRE   = regexp.MustCompile(`^[0-9\x{06F0}-\x{06F9}]+$`)
)

// IsPersianNum reports whether s consists only of Persian digits (if convert==false)
// or allows Persian+ASCII digits (if convert==true). It ignores surrounding
// whitespace but does not accept signs or decimal separators (extend as needed).
func IsPersianNum(s string, convert bool) bool {
	s = irvalidate.TrimSpaces(s)
	if s == "" {
		return false
	}
	if convert {
		return mixedDigitsRE.MatchString(s)
	}
	return persianDigitsRE.MatchString(s)
}

// IsNumericNormalized checks digits after mapping Persian digits to ASCII.
func IsNumericNormalized(s string) bool {
	s = irvalidate.NormalizePersianDigits(irvalidate.TrimSpaces(s))
	return asciiDigitsRE.MatchString(s)
}
