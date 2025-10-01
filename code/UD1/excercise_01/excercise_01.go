package main

import (
	"fmt"
)

func getBinaryAge(age int) string {
	return fmt.Sprintf("%b", age)
}

func countCandles(ageInBinary string) map[string]int {
	candleMap := make(map[string]int)

	candleMap["zero"] = 0
	candleMap["one"] = 0

	for i := 0; i < len(ageInBinary); i++ {
		if ageInBinary[i] == '0' {
			candleMap["zero"] += 1
		} else {
			candleMap["one"] += 1
		}
	}

	return candleMap
}

func displayInformation(age int, ageInBinary string, candles map[string]int, ageInSeconds int64) {
	fmt.Printf("Age: %d, Age in Binary: %b\n", age, age)
	fmt.Printf("Total candles: %d \n", len(ageInBinary))
	fmt.Printf("Unlit: %d Lit: %d \n", candles["zero"], candles["one"])
	fmt.Printf("Lit order: %b \n", age)
	fmt.Printf("Age in seconds: %d \n", ageInSeconds)
}

func getInput(message string) int {
	fmt.Println(message)
	var input int
	fmt.Scanf("%d", &input)
	return input
}

func secondsToAge(ageInSeconds int) int64 {
	return int64(float64(ageInSeconds) / 60 / 60 / 24 / 365)
}

func main() {
	var testCases int = getInput("Introduce the number of test cases")

	for range testCases {
		var ageInput int = getInput("Intoduce your age in seconds")
		ageInBinary := getBinaryAge(ageInput)
		candles := countCandles(ageInBinary)
		displayInformation(int(secondsToAge(ageInput)), ageInBinary, candles, int64(ageInput))
	}

}
