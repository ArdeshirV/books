package mypackage

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func Display(msg string) {
	printMessage(msg)
}

// Returns minimum of number slice and returns zero if the slice is empty
func Min(numbers []float64) (Minimum float64) {
	if numbers == nil || len(numbers) <= 0 {
		Minimum = 0
		return Minimum
	}
	Minimum = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if Minimum > numbers[i] {
			Minimum = numbers[i]
		}
	}
	return Minimum
}

// Returns maximum of number slice and returns zero if the slice is empty
func Max(numbers []float64) (Maximum float64) {
	if numbers == nil || len(numbers) <= 0 {
		Maximum = 0
		return Maximum
	}
	Maximum = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if Maximum < numbers[i] {
			Maximum = numbers[i]
		}
	}
	return Maximum
}
