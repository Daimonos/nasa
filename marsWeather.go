package nasa

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/daimonos/nasa/models"
)

// MarsWeatherURL is the endpoint for NASA's Mars Weather Reporting
const MarsWeatherURL = "https://api.nasa.gov/insight_weather/?api_key=%s&feedtype=json&ver=1.0"

// GetMarsWeather executes the request to get current weather on mars
func GetMarsWeather() (*models.MarsWeatherResp, error) {
	key, err := GetAPIKey()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(MarsWeatherURL, key)
	bytes, reqErr := ExecuteRequest(url)
	if reqErr != nil {
		return nil, reqErr
	}
	var weather map[string]*json.RawMessage
	var solValidity map[string]models.ReadingValidityCheck
	var solWeather map[string]models.MarsWeather
	var solKeys []string
	var solHoursRequired int16
	var solsChecked []string
	var validityMap map[string]*json.RawMessage
	parseErr := json.Unmarshal(bytes, &weather)
	if parseErr != nil {
		return nil, parseErr
	}
	parseErr = json.Unmarshal(*weather["sol_keys"], &solKeys)
	if parseErr != nil {
		return nil, parseErr
	}
	delete(weather, "sol_keys")
	// now we need to unmarshal the sol
	json.Unmarshal(*weather["validity_checks"], &validityMap)
	parseErr = json.Unmarshal(*validityMap["sol_hours_required"], &solHoursRequired)
	if parseErr != nil {
		return nil, parseErr
	}
	parseErr = json.Unmarshal(*validityMap["sols_checked"], &solsChecked)
	if parseErr != nil {
		return nil, parseErr
	}
	delete(validityMap, "sol_hours_required")
	delete(validityMap, "sols_checked")
	*weather["validity_checks"], err = json.Marshal(validityMap)
	if err != nil {
		return nil, err
	}
	parseErr = json.Unmarshal(*weather["validity_checks"], &solValidity)
	if parseErr != nil {
		return nil, parseErr
	}
	delete(weather, "validity_checks")
	// once the sol_keys and validity_checks are parsed and deleted, we should just be
	// left with the sol weather reports
	bytes, err = json.Marshal(weather)
	if err != nil {
		return nil, errors.New("Error converting manipulated JSON back to byte array")
	}
	parseErr = json.Unmarshal(bytes, &solWeather)
	if parseErr != nil {
		return nil, errors.New("Error parsing sol weather")
	}
	// Compile the return struct
	vCheck := models.ValidityChecks{
		SolHoursRequired: solHoursRequired,
		SolsChecked:      solsChecked,
		SolValidity:      solValidity,
	}
	returnStruct := models.MarsWeatherResp{
		SolKeys:        solKeys,
		ValidityChecks: vCheck,
		SolWeather:     solWeather,
	}
	return &returnStruct, nil
}
