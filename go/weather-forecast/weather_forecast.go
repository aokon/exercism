// Package weather contains utilties related to Weather Forecast in the specific location.
package weather

// CurrentCondition keeps information about current weather.
var CurrentCondition string
// CurrentLocation keeps information about location.
var CurrentLocation string

// Forecast fuction prints infomation about weather condition in specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
