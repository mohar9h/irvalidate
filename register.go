package persianvalidation

import (
	"github.com/go-playground/validator/v10"
	"github.com/mohar9h/irvalidate/rule"
)

// RegisterValidators registers all built-in Persian rule on the provided validator.
// It returns an error only if registration of a rule fails.
func RegisterValidators(v *validator.Validate) error {
	if err := v.RegisterValidation("iranian_mobile", rule.ValidateIranianMobile); err != nil {
		return err
	}
	if err := v.RegisterValidation("iranian_national_code", rule.ValidateIranianNationalCode); err != nil {
		return err
	}
	if err := v.RegisterValidation("iranian_postal_code", rule.ValidateIranianPostalCode); err != nil {
		return err
	}
	if err := v.RegisterValidation("iranian_card_number", rule.ValidateIranianCardNumber); err != nil {
		return err
	}
	if err := v.RegisterValidation("persian_text", rule.ValidatePersianText); err != nil {
		return err
	}
	return nil
}
