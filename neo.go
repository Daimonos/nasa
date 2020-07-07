package nasa

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jimdhughes/nasa/models"
)

// NeoAPIUrl is the url to access the NEO api from NASA
const NeoAPIUrl = "https://api.nasa.gov/neo/rest/v1/feed?start_date=%s&end_date=%s&api_key=%s"

// GetFeed returns the feed data from NASA
func GetFeed(startDate, endDate time.Time) (*models.NeoList, error) {
	key, err := GetAPIKey()
	if err != nil {
		return nil, err
	}
	formattedStart, formattedEnd := startDate.Format("2006-01-02"), endDate.Format("2006-01-02")
	url := fmt.Sprintf(NeoAPIUrl, formattedStart, formattedEnd, key)
	bytes, bufferError := ExecuteRequest(url)
	if bufferError != nil {
		return nil, bufferError
	}
	var nl = models.NeoList{}
	jsonErr := json.Unmarshal(bytes, &nl)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &nl, nil
}
