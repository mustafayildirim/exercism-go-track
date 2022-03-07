package lasagna

func PreparationTime(layers []string, averagePreparationTime int) int {

	if averagePreparationTime == 0 {
		averagePreparationTime = 2
	}

	return len(layers) * averagePreparationTime
}

func Quantities(layers []string) (int, float64) {

	sumOfNoodles := 0
	sumOfSauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			sumOfNoodles += 50
		} else if layer == "sauce" {
			sumOfSauce += 0.2
		}

	}
	return sumOfNoodles, sumOfSauce
}

func AddSecretIngredient(friendList []string, myList []string) {
	secretIngredient := friendList[len(friendList)-1]
	myList[len(myList)-1] = secretIngredient
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantities []float64
	calculatedPortion := float64(portions) / 2
	for _, amount := range quantities {
		scaledQuantities = append(scaledQuantities, calculatedPortion*amount)

	}

	return scaledQuantities
}
