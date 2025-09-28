// Package persianvalidation provides Iranian/Persian-specific validation rules
// that integrate with github.com/go-playground/validator/v10.
//
// The package exposes RegisterValidators(v *validator.Validate) to register
// a set of useful rule such as:
// - iranian_mobile
// - iranian_national_code
// - iranian_postal_code
// - iranian_card_number
// - persian_text
//
// Example:
//
//	v := validator.New()
//	persianvalidation.RegisterValidators(v)
package persianvalidation
