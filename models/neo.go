package models

// MissDistance is the estimate of by how much the NEO will miss Earth
type MissDistance struct {
	Astronomical string `json:"astronomical"`
	Kilometers   string `json:"kilometers"`
	Lunar        string `json:"lunar"`
	Miles        string `json:"miles"`
}

// RelativeVelocity is the struct for capturing relative velocity
type RelativeVelocity struct {
	KilometersPerHour   string `json:"kilometers_per_hour"`
	KilometersPerSecond string `json:"kilometers_per_second"`
	MilesPerHour        string `json:"miles_per_hour"`
}

// CloseApproachData is the data that makes up the close approach object
type CloseApproachData struct {
	CloseApproachDate      string           `json:"close_approach_date"`
	CloseApproachDateFull  string           `json:"close_approach_date_full"`
	EpochDateCloseApproach float32          `json:"epoch_date_close_approach"`
	MissDistance           MissDistance     `json:"miss_distance"`
	RelativeVelocity       RelativeVelocity `json:"relative_velocity"`
	OrbitingBody           string           `json:"orbiting_body"`
}

// EstimatedDiameterMeasures is the estimated measurement of the
// diameters
type EstimatedDiameterMeasures struct {
	EstimatedDiameterMin float32 `json:"estimated_diameter_min"`
	EstimatedDiameterMax float32 `json:"estimated_diameter_max"`
}

// EstimatedDiameter is the struct representing the estimated diameters
type EstimatedDiameter struct {
	Kilometers EstimatedDiameterMeasures `json:"kilometers"`
	Meters     EstimatedDiameterMeasures `json:"meters"`
	Miles      EstimatedDiameterMeasures `json:"miles"`
	Feet       EstimatedDiameterMeasures `json:"feet"`
}

// Neo is the man struct for a Near Earth Object
type Neo struct {
	AbsoluteMagnitude              float32             `json:"absolute_magnitude_h"`
	CloseApproachData              []CloseApproachData `json:"close_approach_data"`
	ID                             string              `json:"id"`
	NeoReferenceID                 string              `json:"neo_reference_id"`
	Name                           string              `json:"name"`
	NasaJplURL                     string              `json:"nasa_jpl_url"`
	EstimatedDiameter              EstimatedDiameter   `json:"estimated_diameter"`
	IsPotentiallyHazardousAsteroid bool                `json:"is_potentially_hazardous_asteroid"`
	IsSentryObject                 bool                `json:"is_sentry_object"`
}

// NeoList is the root object representing the result from the api
type NeoList struct {
	ElementCount     int16            `json:"element_count"`
	NearEarthObjects map[string][]Neo `json:"near_earth_objects"`
}
