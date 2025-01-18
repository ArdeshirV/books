 package main

 import (
         "strconv"
 )

func main() {
         for i := 0; i < 1000000; i++ {
                 FizzBuzzV2(1, 100)
         }
}

func FizzBuzzV2(begin, end int) string {
	result := ""
	const Number, Fizz, Buzz, FizzBuzz = 0, 1, 2, 3
	arr := [4]string{"", "Fizz", "Buzz", "FizzBuzz"}

	for i := begin; i <= end; i++ {
		logic := 0
		if i%3 == 0 {
			logic = 1
		}

		if i%5 == 0 {
			logic += 2
		}

		if logic == 0 {
			result += strconv.Itoa(i)
		} else {
			result += arr[logic]
		}
		result += ", "
	}
	return result[:len(result)-2]
}
