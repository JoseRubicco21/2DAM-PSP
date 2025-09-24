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

func getInput() int {
	fmt.Println("Introduce your age.")
	var input int
	fmt.Scanf("%d", &input)
	return input
}

func ageToSeconds(ageInYears int) int64 {
	const (
		secondsPerMinute = 60
		minutesPerHour   = 60
		hoursPerDay      = 24
		daysPerYear      = 365
	)

	secondsPerYear := int64(secondsPerMinute * minutesPerHour * hoursPerDay * daysPerYear)
	return int64(ageInYears) * secondsPerYear
}

func main() {
	var age int = getInput()
	var ageInSeconds = ageToSeconds(age)
	ageInBinary := getBinaryAge(age)

	candles := countCandles(ageInBinary)

	displayInformation(age, ageInBinary, candles, ageInSeconds)

}
