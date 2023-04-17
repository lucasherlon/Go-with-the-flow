package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averageTime int) int {
    totalLayers := len(layers)
    if averageTime == 0 {
        return totalLayers * 2
    }
    return totalLayers * averageTime
}
// TODO: define the 'Quantities()' function
func Quantities(ingredients []string) (int, float64) {
    noodles := 0
    sauce := 0.0
    for _, ingredient := range ingredients {
        if ingredient == "noodles" {
            noodles += 50
        }
    	if ingredient == "sauce" {
            sauce += 0.2
        }
    }
	return noodles, sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
    indexFriend := len(friendList) - 1
    secretIngredient := friendList[indexFriend]
    myIndex := len(myList) - 1
    myList[myIndex] = secretIngredient
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    ratio := float64(portions) / 2.0
    var scaledQuantities []float64
    var temp float64
    for _, quantity := range quantities {
        temp = quantity * ratio
        scaledQuantities = append(scaledQuantities, temp)
    }
	return scaledQuantities
}
