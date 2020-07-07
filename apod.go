package nasa

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jimdhughes/nasa/models"
)

// ApodURL is the url of NASA's APOD
const ApodURL = "https://api.nasa.gov/planetary/apod?date=%s&hd=%s&api_key=%s"

// GetApod returns NASA's Picture of the day for a given date
func GetApod(date time.Time, hd bool) (*models.Apod, error) {
	key, err := GetAPIKey()
	if err != nil {
		return nil, err
	}
	dateFormat := date.Format("2006-01-02")
	var hdStr string
	hdStr = "false"
	if hd == true {
		hdStr = "true'"
	}
	url := fmt.Sprintf(ApodURL, dateFormat, hdStr, key)
	bytes, reqError := ExecuteRequest(url)
	if reqError != nil {
		return nil, reqError
	}
	var apod models.Apod
	jsonErr := json.Unmarshal(bytes, &apod)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &apod, nil
}
