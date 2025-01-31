package mypackage

import (
	"fmt"
	"errors"
)

func printMessage(message string) {
	fmt.Println(message)
}

func Display(msg string) {
	printMessage(msg)
}

func Average(args []float64) (float64, error) {
	if args == nil {
		return 0.0, errors.New("The Average(args []float64) func failed with nil arguments")
	}
	if len(args) <= 0 {
		return 0.0, nil
	}
	res := 0.0
	for _, f := range args {
		res += f
	}
	return res / float64(len(args)), nil
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
