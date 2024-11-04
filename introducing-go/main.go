package main

// book - Introduce Go

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

type binary int

const copyright = "Copyright(C) 2024 github.com/ArdeshirV, Licensed Under GPLv3+"

func main() {
	//ChapterOne()
	//ChapterTwo()
	//ChapterThree()
	//ChapterFour()
	//ChapterFive()
	//ChapterFivePart2()
	ChapterFivePart3()
}

func ChapterFivePart3() {
	fmt.Println("Chapter five-part3 (slices):")

	slice1 := make([]int, 10, 100)
	for i, value := range slice1 {
		slice1[i] = value * 3
	}
	slice1[9] = 99
	slice1 = append(slice1, slice1...)
	slice1[10] = 100
	//slice2 := make([]int, len(slice1))
	//copy(slice2, slice1)
	slice2 := slice1[:]
	slice1 = slice2
	for i, value := range slice1 {
		fmt.Printf("[%d]=%d ", i, value)
	}
	fmt.Print("\n\n")
	slice1[4] = 33
	for i, value := range slice2 {
		fmt.Printf("[%d]=%d ", i, value)
	}
	fmt.Println()

}

func ChapterFivePart2() {
	fmt.Println("Chapter five-part2:")
	const n = 10
	//fmt.Println(PerformFunc(fibonacci_1, n))
	//fmt.Println(PerformFunc(fibonacci_2, n))
	//fmt.Println(MyToUpperCase("Hello, World!"))
	arr := []int{50, 10, 60, 30, 40, 1, 20}
	arr = MergeSort(arr)

	var sb strings.Builder
	for i := 0; i < len(arr); i++ {
		//sb.WriteString(fmt.Sprintf("arr[%v]=%v, ", i, arr[i]))
		sb.WriteString("arr[")
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString("]=")
		sb.WriteString(strconv.Itoa(arr[i]))
		sb.WriteString(", ")
	}
	output, _ := strings.CutSuffix(sb.String(), ", ")

	target := 50
	index := BinarySearch(arr, target)
	output += fmt.Sprintf("\nindex of %v is %v\n", target, index)
	fmt.Print(output)

	fmt.Println(seniorToUpper("Hello, World!"))

	end := unicode.Avestan.R32[0].Hi
	begin := unicode.Avestan.R32[0].Lo
	for i := begin; i < end; i++ {
		fmt.Print(string(rune(i)))
	}
	fmt.Println()

	fmt.Println(binary('A'))
	fmt.Println(binary('a'))
}

func (b binary) String() string {
	return fmt.Sprintf("%b", b)
}

func seniorToUpper(text string) string {
	res := make([]rune, len(text))
	for i, r := range text {
		if r >= 'a' && r <= 'z' {
			res[i] = r | 32
		} else {
			res[i] = r
		}
	}
	return string(res)
}

func juniorToUpper2(text string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r - ('a' - 'A')
		} else {
			return r
		}
	}, text)
}

func juniorToUpper(text string) string {
	res := make([]rune, len(text))
	for i, r := range text {
		if r >= 'a' && r <= 'z' {
			res[i] = r + ('a' - 'A')
		} else {
			res[i] = r
		}
	}
	return string(res)
}

func Factorial(n int) (res int) {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(a []int, b []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res
}

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func PerformFunc(fib func(int) int, n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += fmt.Sprintf("%v, ", fib(i))
	}
	res = res[:len(res)-2]
	return res
}

func fibonacci_1(n int) int {
	if n < 2 {
		return 1
	}
	return fibonacci_1(n-1) + fibonacci_1(n-2)
}

func fibonacci_2(n int) int {
	if n < 2 {
		return 1
	}

	a, b, res := 0, 1, 1
	for i := 1; i < n; i++ {
		a = b
		b = res
		res = a + b
	}
	return res
}

func fibonacci_3(n int) int {
	return 0
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
	key := "7"
	res := encode_xor(key, myName)
	fmt.Println("<", res, ">")
	res = encode_xor(key, res)
	fmt.Println("<", res, ">")
	myNameBytes := []byte(myName)
	fmt.Printf("%q\n", myNameBytes)

	fmt.Println("Hello One")
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			fmt.Println("Hello Inside")
			for counter := 0; counter < 10; counter++ {
				fmt.Printf("%v, ", counter)
			}
			fmt.Println()
			wg.Done()
		}()
	}
	fmt.Println("Hello Two")
	for i := 0; i < 10; i++ {
		wg.Wait()
	}
	fmt.Println("Hello Three")
	// ...
	res = fmt.Sprintf("%d + %d = %d", 10, 20, 30)
	fmt.Println(res)
	var i int
	i = 9
	fmt.Println("i=", i)
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
		fallthrough
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

type Stack struct {
	items []int
}

func SeniorToUpper(text string) string {
	res := ""

	return res
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (item int) {
	if len(s.items) == 0 {
		return -1
	}
	item = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func ChapterOne() {
	fmt.Println("Chapter one:")
	res, err := fmt.Printf("The Golang Programming Language\n")
	fmt.Printf("res = %v, err = %v\n", res, err)
	os.Exit(0)
	fmt.Println("You won't see this line never")
}
