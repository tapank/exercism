// Package weather provides a single function that stringifies
// weather forecast for a city.
package weather

// CurrentCondition is the weather forecast for a city.
var CurrentCondition string

// CurrentLocation is the city name.
var CurrentLocation string

// Forecast creates a human readable string for given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
