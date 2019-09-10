package models

// Reading is the format of the JSON response for Mars Weather
type Reading struct {
	Average float32 `json:"av"`
	Current float32 `json:"ct"`
	Max     float32 `json:"mx"`
	Min     float32 `json:"mn"`
}

type WDReading struct {
	CompassDegress float32 `json:"compass_degrees"`
	CompassPoint string `json:"compass_point"`
	CompassRight float32 `json:"compass_right"`
	CompassUp float32 `json:"compass_up"`
	Samplecount int16 `json:"ct"`
}

// ValidityCheck is the valididty check
type ValidityCheck struct {
	SolHoursWithData []int8 `json:"sol_hours_with_data"`
	Valid            bool   `json:"valid"`
}

// ReadingValidityCheck is the ReadingValidityCheck
type ReadingValidityCheck struct {
	AtmosphericValidity         ValidityCheck `json:"AT"`
	AtmosphericPressureValidity ValidityCheck `json:"PRE"`
	WindDirectionValidity       ValidityCheck `json:"WD"`
	HWS                         ValidityCheck `json:"HWS"`
}

// MarsWeather is the format of a reading at a SOL
type MarsWeather struct {
	AtmosphericTemperature Reading            `json:"AT"`
	AtmosphericPressure    Reading            `json:"PRE"`
	WindDirection          map[string]WDReading `json:"WD"`
	HWS                    Reading            `json:"HWS"`
	FirstUTC               string             `json:"First_UTC"`
	LastUTC                string             `json:"Last_UTC"`
	Season                 string             `json:"season"`
}

// ValidityChecks are all the validity checks
type ValidityChecks struct {
	SolHoursRequired int16                           `json:"sol_hours_required"`
	SolsChecked      []string                        `json:"sols_checked"`
	SolValidity      map[string]ReadingValidityCheck  `json:"sol_validity_checks"`
}

// MarsWeatherResp is a map of Weather measured at a SOL identified by the map string
type MarsWeatherResp struct {
	SolKeys        []string               `json:"sol_keys"`
	ValidityChecks ValidityChecks          `json:"validity_checks"`
	SolWeather     map[string]MarsWeather `json:"sol_weather"`
}


