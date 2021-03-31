package models

type MarsRoverPhoto struct {
	ID int64 `json:"id"`
	Sol int64 `json:"sol"`
	Camera Camera `json:"camera"`
	ImageSource string `json:"img_src"`
	EarthDate string `json:"earth_date"`
	Rover Rover `json:"rover"`
}

type MarsRoverPhotosResponse struct {
	Photos []MarsRoverPhoto `json:"photos"`
}

type MarsRoverPhotosQueryParameters struct {
	Sol string `json:"sol"`
	Camera string `json:"camera"`
	Page string `json:"page"`
	EarthDate string `json:"earth_date"`
}