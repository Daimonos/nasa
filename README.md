# NASA

## Description

This is a simple library wrapping some of the NASA API endpoints.
Can be used with Go applications.
See the API docs at https://api.nasa.gov

## Usage

Requires an environment variable to be set `NASA_API_KEY`

NASA has a public API Key `DEMO_KEY` which has some limitations. Signing up for an API Key is easy and fast

## Notes & Warnings

### Mars Weather Endpoints

I really didn't like the way the Mars Weather endpoint data was coming back.

In order to unmarshal and access better, I've put all the 'Weather' reports into a new key called 'sol_weather' and the validation checks for each of the sols into a new key called 'sol_validity_check'. Please confirm usage of this wrapper vs the official docs to see differences
