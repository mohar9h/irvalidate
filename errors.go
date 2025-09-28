package irvalidate

import "fmt"

// ValidationError is a lightweight wrapper used when callers want
// to return clearer errors from helper functions.
type ValidationError struct {
	Field string
	Rule  string
	Msg   string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s (%s)", e.Field, e.Msg, e.Rule)
}
