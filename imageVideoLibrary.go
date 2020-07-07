package nasa

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jimdhughes/nasa/models"
)

// BaseMarsRoverImageURL is the base URL endpoint for the Mars Rover Image API
const BaseMarsRoverImageURL = "https://images-api.nasa.gov"

// MarsRoverSearchURL is the url for general searching of the rover image api
const MarsRoverSearchURL = BaseMarsRoverImageURL + "/search?api_key=%s"

// MarsRoverAssetURL is the url for searching for specific asset images
const MarsRoverAssetURL = BaseMarsRoverImageURL + "/asset/%s?api_key=%s"

// MarsRoverMetadataURL is the url for searching for rover metadata
const MarsRoverMetadataURL = BaseMarsRoverImageURL + "/metadata/%s?api_key=%s"

// MarsRoverVideoCaptionsURL is the endpoint for vido captions
const MarsRoverVideoCaptionsURL = BaseMarsRoverImageURL + "/captions/%s?api_key=%s"

// MarsRoverAlbumContentURL is the endpoint for album content
const MarsRoverAlbumContentURL = BaseMarsRoverImageURL + "/album/%s?api_key=%s"

// MarsRoverImageSearchQuery is the query object for the SearchImage API Endpoint
type MarsRoverImageSearchQuery struct {
	Q                string `json:"q"`
	Center           string `json:"center"`
	Description      string `json:"description"`
	Description508   string `json:"description508"`
	Keywords         string `json:"keywords"`
	Location         string `json:"location"`
	MediaType        string `json:"mediaType"`
	NasaID           string `json:"nasaID"`
	Page             string `json:"page"`
	Photographer     string `json:"photographer"`
	SecondaryCreator string `json:"secondaryCreator"`
	Title            string `json:"title"`
	YearStart        string `json:"yearStart"`
	YearEnd          string `json:"yearEnd"`
}

// SearchMarsRoverImages is an API wrapper to fetch images from the
// Mars Rover API. All parameters are optional and should be a blank string
// if not set
func SearchMarsRoverImages(query MarsRoverImageSearchQuery) (*models.SearchResponse, error) {
	key, err := GetAPIKey()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(MarsRoverSearchURL, key)
	if query.Q != "" {
		url += "&q=%s"
		url = fmt.Sprintf(url, query.Q)
	}
	if query.Center != "" {
		url += "&center=%s"
		url = fmt.Sprintf(url, query.Center)
	}
	if query.Description != "" {
		url += "&description=%s"
		url = fmt.Sprintf(url, query.Description)
	}
	if query.Description508 != "" {
		url += "description_508=%s"
		url = fmt.Sprintf(url, query.Description508)
	}
	if query.Keywords != "" {
		url += "&keywords=%s"
		url = fmt.Sprintf(url, query.Keywords)
	}
	if query.Location != "" {
		url += "&location=%s"
		url = fmt.Sprintf(url, query.Location)
	}
	if query.MediaType != "" {
		url += "&media_type=%s"
		url = fmt.Sprintf(url, query.MediaType)
	}
	if query.NasaID != "" {
		url += "&nasa_id=%s"
		url = fmt.Sprintf(url, query.NasaID)
	}
	if query.Page != "" {
		url += "&page=%s"
		url = fmt.Sprintf(url, query.Page)
	}
	if query.Photographer != "" {
		url += "&photographer=%s"
		url = fmt.Sprintf(url, query.Photographer)
	}
	if query.SecondaryCreator != "" {
		url += "&secondary_creator=%s"
		url = fmt.Sprintf(url, query.SecondaryCreator)
	}
	if query.Title != "" {
		url += "&title=%s"
		url = fmt.Sprintf(url, query.Title)
	}
	if query.YearStart != "" {
		url += "&year_start=%s"
		url = fmt.Sprintf(url, query.YearStart)
	}
	if query.YearEnd != "" {
		url += "&year_end=%s"
		url = fmt.Sprintf(url, query.YearEnd)
	}
	bytes, err := ExecuteRequest(url)
	if err != nil {
		return nil, err
	}
	var result models.SearchResponse
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAssetMediaManifest returns the media manifest for a given nasaID
func GetAssetMediaManifest(nasaID string) (*models.AssetResponse, error) {
	key, err := GetAPIKey()
	if err != nil {
		return nil, err
	}
	if nasaID == "" {
		return nil, errors.New("Require non empty NASA ID")
	}
	url := fmt.Sprintf(MarsRoverAssetURL, nasaID, key)

	bytes, err := ExecuteRequest(url)
	if err != nil {
		return nil, err
	}
	var result models.AssetResponse
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAssetMetadataManifest returns the metadata for a specific nasaID
func GetAssetMetadataManifest(nasaID string) (*models.MetadataResponse, error) {
	key, err := GetAPIKey()
	if err != nil {
		return nil, err
	}
	if nasaID == "" {
		return nil, errors.New("Require non empty NASA ID")
	}
	url := fmt.Sprintf(MarsRoverMetadataURL, nasaID, key)

	bytes, err := ExecuteRequest(url)
	if err != nil {
		return nil, err
	}
	var result models.MetadataResponse
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMediaVideoCaptionsLocation gets the location of video captions for a specific nasaID
func GetMediaVideoCaptionsLocation(nasaID string) (*models.VideoCaptionResponse, error) {
	key, err := GetAPIKey()
	if err != nil {
		return nil, err
	}
	if nasaID == "" {
		return nil, errors.New("Require non empty NASA ID")
	}
	url := fmt.Sprintf(MarsRoverVideoCaptionsURL, nasaID, key)
	bytes, err := ExecuteRequest(url)
	if err != nil {
		return nil, err
	}
	var result models.VideoCaptionResponse
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMediaAlbumContents returns the contents of a specific albumName. page is optional
func GetMediaAlbumContents(albumName string, page int) (*models.AlbumResponse, error) {
	key, err := GetAPIKey()
	if err != nil {
		return nil, err
	}
	if albumName == "" {
		return nil, errors.New("Require non empty NASA ID")
	}
	url := fmt.Sprintf(MarsRoverAlbumContentURL, albumName, key)
	if page > 0 {
		url += "&page=%d"
		url = fmt.Sprintf(url, page)
	}
	bytes, err := ExecuteRequest(url)
	if err != nil {
		return nil, err
	}
	var result models.AlbumResponse
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
