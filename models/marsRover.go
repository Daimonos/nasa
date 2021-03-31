package models

type Rover struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	LandingDate string `json:"landing_date"`
	LaunchDate string `json:"launch_date"`
	Status string `json:"status"`
	MaxSol string `json:"max_sol"`
	MaxDate string `json:"max_date"`
	TotalPhotos int64 `json:"total_photos"`
	Cameras []Camera `json:"cameras"`
}

type Camera struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	RoverID string `json:"rover_id"`
	FullName string `json:"full_name"`
}

type ListRoversResponse struct {
	Rovers []Rover `json:"rovers"`
}

type GetRoverResponse struct {
	Rover Rover `json:"rover"`
}