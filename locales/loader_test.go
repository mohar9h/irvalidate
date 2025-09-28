package locales

import (
	"os"
	"testing"
)

func setupTestLocaleFile(t *testing.T, filename, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
}

func teardownTestLocaleFile(filename string) {
	_ = os.Remove(filename)
}

func TestLoadMessages_AndGetMessage(t *testing.T) {
	filename := "locales/test.json"
	content := `{"hello":"سلام","bye":"خداحافظ"}`
	_ = os.MkdirAll("locales", 0755)
	setupTestLocaleFile(t, filename, content)
	defer teardownTestLocaleFile(filename)

	err := LoadMessages("test")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	val, err := GetMessage("test", "hello")
	if err != nil {
		t.Fatalf("expected key to exist, got error: %v", err)
	}

	if val.(string) != "سلام" {
		t.Errorf("expected سلام, got %v", val)
	}
}

func TestGetMessage_Fallback(t *testing.T) {
	Messages["en"] = map[string]interface{}{
		"fallback_key": "English message",
	}

	val, err := GetMessage("fa", "fallback_key")
	if err != nil {
		t.Fatalf("expected fallback to work, got error: %v", err)
	}

	if val.(string) != "English message" {
		t.Errorf("expected English message, got %v", val)
	}
}

func TestGetMessage_NotFound(t *testing.T) {
	Messages = map[string]map[string]interface{}{} // reset

	_, err := GetMessage("fa", "missing_key")
	if err == nil {
		t.Error("expected error for missing key, got nil")
	}
}
