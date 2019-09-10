package nasa

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

// GetAPIKey is a utility function to get the API Key for NASA
func GetAPIKey() (string, error) {
	key := os.Getenv("NASA_API_KEY")
	if key == "" {
		return "", errors.New("NASA_API_KEY must be set")
	}
	return key, nil
}

// ExecuteRequest is a helper function
func ExecuteRequest(url string) ([]byte, error) {
	resp, respErr := http.Get(url)
	if respErr != nil {
		return nil, respErr
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Received Non-200 Status Code from NASA API")
	}
	bytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}
	return bytes, nil
}
