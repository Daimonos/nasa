package nasa

import (
	"os"
	"testing"
)

func TestNoApiKey(t *testing.T) {
	key, _ := GetAPIKey()
	if key != "" {
		t.Errorf("Expected no API Key, got %s", key)
	}
}

func TestApiKey(t *testing.T) {
	os.Setenv("NASA_API_KEY", "AKEY")
	key, _ := GetAPIKey()
	if key != "AKEY" {
		t.Errorf("Expected key 'API_KEY' got: %s", key)
	}
}
