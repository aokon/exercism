package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var result bool

	if kind == "car" || kind == "truck" {
		result = true
	} else {
		result = false
	}

	return result
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var betterOption string

	if strings.Compare(option1, option2) == -1 {
		betterOption = option1
	} else {
		betterOption = option2
	}

	return fmt.Sprintf("%s is clearly the better choice.", betterOption)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var armotisationFactor float64

	if age >= 1 && age < 3 {
		armotisationFactor = 0.2
	} else if age >= 3 && age < 10 {
		armotisationFactor = 0.3
	} else {
		armotisationFactor = 0.5
	}

	return (1 - armotisationFactor) * originalPrice
}
