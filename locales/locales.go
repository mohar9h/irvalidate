package locales

import (
	"encoding/json"
	"os"
)

var Messages map[string]string

func LoadMessages(locale string) error {
	data, err := os.ReadFile("locales/" + locale + ".json")
	if err != nil {
		data, _ = os.ReadFile("locales/en.json")
	}

	return json.Unmarshal(data, &Messages)
}
