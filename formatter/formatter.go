package formatter

import (
	"fmt"
	"strings"
)

// FormatMessage replaces placeholders like :attribute, :example, :days
// Supports single values and list placeholders ([]string).
func FormatMessage(template string, data map[string]interface{}, separator string) string {
	result := template

	for key, val := range data {
		placeholder := ":" + key

		switch v := val.(type) {
		case string:
			result = strings.ReplaceAll(result, placeholder, v)
		case []string:
			result = strings.ReplaceAll(result, placeholder, strings.Join(v, separator))
		default:
			result = strings.ReplaceAll(result, placeholder, fmt.Sprintf("%v", v))
		}
	}

	return result
}
