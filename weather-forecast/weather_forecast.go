// Package weather provides a library for obtaining weather forecast information.
package weather

// CurrentCondition is a global variable to store the current weather condition.
var CurrentCondition string

// CurrentLocation is a global variable to store the current location.
var CurrentLocation string

// Forecast returns the current weather condition for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
