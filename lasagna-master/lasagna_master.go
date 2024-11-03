package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0

	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		default:
			// not noodles or sauce, ignore
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledPortions := make([]float64, len(quantities))

	for i := 0; i < len(quantities); i++ {
		scaledPortions[i] = quantities[i] * float64(portions) / 2
	}
	return scaledPortions
}
