package main

import (
	"fmt"
	"math"
	"os"
)

const copyright = "Copyright(C) 2024 github.com/ArdeshirV, Licensed Under GPLv3+"

func main() {
	//ChapterOne()
	//ChapterTwo()
	//ChapterThree()
	//ChapterFour()
	ChapterFive()
}

func ChapterFive() {
	fmt.Println("Chapter five:")

	// Arrays
	fmt.Println("Arrays")
	var x [5]int
	x[4] = 100
	fmt.Printf("%v\n", x)

	// Average by arrays
	y := [5]float64{98, 77, 94, 65, 95}
	var total float64 = 0.0
	for _, value := range y {
		total += value
	}
	fmt.Println(total / float64(len(y)))

	fmt.Println("\nSlices:")
	myName := "Ardeshir"
	key := "1"
	res := encode_xor(key, myName)
	fmt.Println("<", res, ">")
	res = encode_xor(key, res)
	fmt.Println("<", res, ">")
	myNameBytes := []byte(myName)
	fmt.Printf("%q\n", myNameBytes)
	// ...

}

func encode_xor(key, text string) (res string) {
	keyByte := byte(key[0])
	tempBytes := []byte(text)
	for i, _ := range tempBytes {
		tempBytes[i] ^= keyByte
	}
	res = string(tempBytes)
	return res
}

func ChapterFour() {
	fmt.Println("Chapter four:")
	for i := 1; i <= 10; i++ {
		var numberStatus string
		if i%2 == 0 {
			numberStatus = "even"
		} else {
			numberStatus = "odd"
		}
		fmt.Println(i, numberStatus)
	}

	fmt.Print("Enter an integer number between 0-9: ")
	var number int
	var numberValue string
	// fmt.Scanf("%d", &number)
	number = 3
	switch number {
	case 0:
		numberValue = "Zero"
	case 1:
		numberValue = "One"
	case 2:
		numberValue = "Two"
	case 3:
		numberValue = "Three"
	case 4:
		numberValue = "Four"
	case 5:
		numberValue = "Five"
	case 6:
		numberValue = "Six"
	case 7:
		numberValue = "Seven"
	case 8:
		numberValue = "Eight"
	case 9:
		numberValue = "Nine"
	default:
		numberValue = "Uknown Numer"
	}
	arrNumbers := [10]string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	fmt.Println(number, " = ", numberValue)
	fmt.Println(number, " = ", arrNumbers[number])

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(i, ", ")
		}
	}
	fmt.Println("\b\b  ")

	const seperator = ", "
	output := ""
	arrStates := [4]string{"", "Fizz", "Buzz", "FizzBuzz"}
	for i := 1; i <= 100; i++ {
		state := FizzBuzzLogic(i)
		if state == 0 {
			arrStates[0] = fmt.Sprintf("%d", i)
		}
		output += arrStates[state] + seperator
	}
	output = output[:len(output)-2]
	fmt.Println(output)
}

func FizzBuzzLogic(number int) (res int) {
	res = 0
	if number%3 == 0 {
		res = 1
	}
	if number%5 == 0 {
		res |= 2
	}
	return res
}

func ChapterThree() {
	fmt.Println("Chapter three:")
	fmt.Println("math.Pi =", math.Pi)
	fmt.Println(copyright)
	fmt.Print("Enter temperature in fahrenheit: ")

	var fahrenheit float64
	if _, err := fmt.Scanf("%f", &fahrenheit); err != nil {
		panic(err)
	}
	celcious := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Temperature in degree is %v\n", celcious)
	fmt.Println()

	var feet float64
	fmt.Print("Enter the height in feet: ")
	if _, err := fmt.Scanf("%f", &feet); err != nil {
		panic(err)
	}
	meter := feet * 0.3048
	fmt.Printf("The height in meter is %v\n", meter)
}

func ChapterTwo() {
	fmt.Println("Chapter two:")
	a, b, c := 1.1, 2., 0.
	a = 0 / (a - 1.1)
	c = a + b
	fmt.Printf("%v + %v = %v\n", a, b, c)
	var some_string string
	some_string = "Hello, World!"
	fmt.Println(some_string)
	for i := 0; i < len(some_string); i++ {
		fmt.Printf("%c ", some_string[i])
	}
	fmt.Println()
	var OutputMessage string
	OutputMessage = fmt.Sprintf("some_string[1] = %c\n", some_string[1])
	fmt.Print(OutputMessage)
	var cmp1 complex128
	cmp2 := 10i + 13
	cmp1 = cmp2
	fmt.Println("cmp1 =", cmp1)
}

func ChapterOne() {
	fmt.Println("Chapter one:")
	res, err := fmt.Printf("The Golang Programming Language\n")
	fmt.Printf("res = %v, err = %v\n", res, err)
	os.Exit(0)
	fmt.Println("You won't see this line never")
}
