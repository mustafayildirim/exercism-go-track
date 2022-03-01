package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {

	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	car := option1
	if option1 > option2 {
		car = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", car)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

	defaultCostDiscountPercentage := 0.70
	if age < 3 {
		defaultCostDiscountPercentage = 0.80
	} else if age > 9 {
		defaultCostDiscountPercentage = 0.50
	}
	return originalPrice * defaultCostDiscountPercentage
}
