package locales

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var Messages = make(map[string]map[string]interface{}) // locale → key → value

// LoadMessages loads messages for a given locale (fa, en, etc.)
func LoadMessages(locale string) error {
	file := filepath.Join("locales", locale+".json")
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	var msgs map[string]interface{}
	if err := json.Unmarshal(data, &msgs); err != nil {
		return err
	}

	Messages[locale] = msgs
	return nil
}

// GetMessage fetches a message template, falling back if needed
func GetMessage(locale, key string) (interface{}, error) {
	if msgs, ok := Messages[locale]; ok {
		if val, exists := msgs[key]; exists {
			return val, nil
		}
	}

	// fallback to English
	if msgs, ok := Messages["en"]; ok {
		if val, exists := msgs[key]; exists {
			return val, nil
		}
	}

	return nil, errors.New("message not found" + key)
}
