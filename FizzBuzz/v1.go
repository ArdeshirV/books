package main

import (
	"strconv"
)

func main() {
	for i := 0; i < 1000000; i++ {
		FizzBuzzV1(1, 100)
	}
}

func FizzBuzzV1(begin, end int) string {
	result := ""
	for i := begin; i <= end; i++ {
		if i%15 == 0 {
			result += "FizzBuzz"
		} else if i%5 == 0 {
			result += "Buzz"
		} else if i%3 == 0 {
			result += "Fizz"
		} else {
			result += strconv.Itoa(i)
		}
		result += ", "
	}
	return result[:len(result)-2]
}
