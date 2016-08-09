package main

import "fmt"

var availableCoins = [7]int{100, 50, 25, 20, 10, 5, 1}

func main() {
	fmt.Printf("available coins: %v\n", availableCoins)
	overallValue := 197
	var finalCoins []int
	var counter int
	finalCoins = makeChange(overallValue, counter, finalCoins)
	fmt.Printf("final coins: %v\n", finalCoins)
}

func makeChange(overallValue int, counter int, finalCoins []int) []int {

	currentCoin := availableCoins[counter]

	if overallValue >= currentCoin {

		if counter < len(availableCoins)-1 &&
			(overallValue%availableCoins[counter+1] < overallValue%currentCoin) &&
			availableCoins[counter+1]*2 > currentCoin {
			currentCoin = availableCoins[counter+1]
			counter++
		}

		fmt.Printf("Adding %v coin to finalCoins\n", currentCoin)
		finalCoins = append(finalCoins, currentCoin)
		overallValue -= currentCoin
	} else {
		counter++
	}

	if overallValue > 0 {
		return makeChange(overallValue, counter, finalCoins)
	}

	return finalCoins
}
