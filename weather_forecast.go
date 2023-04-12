// Package weather provides a function that forecasts the weather conditions
// at several cities in Goblinocus.
package weather

// CurrentCondition is a variable that brings the current weather condition in a given city.
var CurrentCondition string
// CurrentLocation is a variable that brings a city that we want to forecast the weather.
var CurrentLocation string

// Forecast returns a string which is the weather forecast for a city in Goblinocus.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
