package formatter

import "testing"

func TestFormatMessage_Simple(t *testing.T) {
	tmpl := "The :attribute must be at least :min characters."
	data := map[string]interface{}{
		"attribute": "Password",
		"min":       "8",
	}

	got := FormatMessage(tmpl, data, ", ")
	want := "The Password must be at least 8 characters."

	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestFormatMessage_List(t *testing.T) {
	tmpl := ":attribute must be one of (:days)."
	data := map[string]interface{}{
		"attribute": "Day",
		"days":      []string{"Sat", "Sun", "Mon"},
	}

	got := FormatMessage(tmpl, data, ", ")
	want := "Day must be one of (Sat, Sun, Mon)."

	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestFormatMessage_UnknownType(t *testing.T) {
	tmpl := "Value: :val"
	data := map[string]interface{}{
		"val": 12345, // int instead of string
	}

	got := FormatMessage(tmpl, data, ", ")
	want := "Value: 12345"

	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
