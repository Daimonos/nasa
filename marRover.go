package nasa

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jimdhughes/nasa/models"
)

const MarsRoverBaseURL = "https://api.nasa.gov/mars-photos/api/v1/rovers?API_KEY=%s"
const MarsRoverDetailURL = "https://api.nasa.gov/mars-photos/api/v1/rovers/%s?API_KEY=%s"
const MarsRoverPhotosURL="https://api.nasa.gov/mars-photos/api/v1/rovers/%s?%s&API_KEY=%s"

func ListMarsRovers() ([]models.Rover, error) {
	var roversResponse models.ListRoversResponse
	key, err := GetAPIKey()
	if err != nil {
		return nil, errors.New("Error getting API Key")
	}
	url := fmt.Sprintf(MarsRoverBaseURL, key)
	bytes, reqErr := ExecuteRequest(url)
	if reqErr != nil {
		return nil, errors.New("Error executing request - " +reqErr.Error())
	}
	parseErr := json.Unmarshal(bytes, &roversResponse)
	if parseErr != nil {
		return nil, errors.New("Error parsing rovers response - " + parseErr.Error())
	}
	return roversResponse.Rovers, nil
}

func GetMarsRover(rover string) (models.Rover, error) {
	var roverResponse models.GetRoverResponse
	key, err := GetAPIKey()
	if err != nil {
		return models.Rover{}, errors.New("Error getting API Key")
	}
	url := fmt.Sprintf(MarsRoverBaseURL, key)
	bytes, reqErr := ExecuteRequest(url)
	if reqErr != nil {
		return models.Rover{}, errors.New("Error executing request - " +reqErr.Error())
	}
	parseErr := json.Unmarshal(bytes, &roverResponse)
	if parseErr != nil {
		return models.Rover{}, errors.New("Error parsing rover response - " + parseErr.Error())
	}
	return roverResponse.Rover, nil
}

func GetMarsRoverPhotos(rover string, query models.MarsRoverPhotosQueryParameters) ([]models.MarsRoverPhoto, error) {
	if rover == "" {
		return nil, errors.New("Please provide a valid rover name")
	}
	if query.Sol != "" && query.EarthDate != "" {
		return nil, errors.New("Please only request data for a Sol or an Earth Date, not both")
	}
	var photoResponse models.MarsRoverPhotosResponse
		key, err := GetAPIKey()
	if err != nil {
		return nil, errors.New("Error getting API Key")
	}
	queryString := extractQuerystringFromSearchQuery(query)
	url := fmt.Sprintf(MarsRoverPhotosURL, rover, queryString, key)
	bytes, reqErr := ExecuteRequest(url)
	if reqErr != nil {
		return nil, errors.New("Error executing request - " +reqErr.Error())
	}
	parseErr := json.Unmarshal(bytes, &photoResponse)
	if parseErr != nil {
		return nil, errors.New("Error parsing rover response - " + parseErr.Error())
	}
	return photoResponse.Photos, nil
}


func extractQuerystringFromSearchQuery(query models.MarsRoverPhotosQueryParameters) string {
	queryString := ""
	if query.Camera != "" {
		queryString+="&camera="+query.Camera
	}
	if query.EarthDate != "" {
		queryString += "&earth_date="+query.EarthDate
	}
	if query.Sol != "" {
		queryString+="&sol="+query.Sol
	}
	if query.Page != "" {
		queryString+="&page="+query.Sol
	}
	if len(queryString) > 0 {
		queryString=queryString[1:]
	}
	return queryString
}
