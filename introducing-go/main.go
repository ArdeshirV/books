package main

// book - Introduce Go

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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
	//ChapterFivePart3()
	//ChapterFivePart4()
	//FizzBuzz(1, 30)
	//FizzBuzzV2(1, 30)
	//Chapter5()
	Chapter6()
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() uint {
		i += 2
		return i
	}
}

func Chapter6() {
	nextEven := makeEvenGenerator()
	for i := 0; i < 10; i++ {
		fmt.Print(nextEven(), " ")
	}
	fmt.Println()

	//
}

func Chapter5() {
	var x [5]int
	x[2] = 3
	fmt.Println(x)

	y := make([]float64, 10, 100)
	for i := 0; i < 200; i++ {
		y = append(y, float64(i))
	}
	for _, value := range y {
		fmt.Print(value, ", ")
	}
	fmt.Println()

	arr1 := []int{1, 2, 3}
	arr2 := []int{10, 10, 10}
	arr2[1] = 999
	copy(arr2, arr1)
	arr2[0] = 200
	arr2[2] = -50
	fmt.Println(arr1)

	m := make(map[string]int, 2)
	m["one"] = 1
	m["two"] = 2
	delete(m, "one")
	fmt.Println(m)
	key := "two"
	if value, ok := m[key]; ok {
		fmt.Printf("[%v] = %v\n", key, value)
	}
}

func FizzBuzzV2(begin, end int) {
	fmt.Printf("FizzBuzz from %v to %v:\n", begin, end)

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
	result = result[:len(result)-2]
	fmt.Println(result)
}

func FizzBuzz(begin, end int) {
	fmt.Printf("FizzBuzz from %v to %v:", begin, end)
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
	result = result[:len(result)-2]
	fmt.Println(result)
}

func ChapterFivePart4() {
	fmt.Println("\033[1;27mChapter five-part3 (slices):\033[0m")
	fmt.Println(argumentEvaluator("Ardeshir"))
	fmt.Println("Hello World")
	fmt.Println(100 / 3)
	fmt.Println(strconv.Itoa((int)("Hello, World"[3])))
	fmt.Println(0b11111111)
	var x, y int
	x, y = 10, 20
	x += y
	fmt.Println(x, y)

	fmt.Println("Enter fahrenheit: ")
	var fahrenheit float64
	//fmt.Scanf("%f", &fahrenheit)
	fahrenheit = 123
	const mul float64 = 5.0 / 9.0
	centigrade := (fahrenheit - 32) * mul
	fmt.Printf("Celsius: %.2f\n", centigrade)

	const f2m float64 = 0.3048
	var feet, meters float64
	feet = 170.0
	meters = feet * f2m
	fmt.Printf("meters: %.3f\n", meters)

	i := 0
	for i < 10 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	for i, j := 10, 20; i < j; i, j = i+1, j-1 {
		fmt.Print(i, j, " - ")
	}
	fmt.Println()

	i = 99
	i--
	if i%2 == 0 {
		fmt.Printf("%v %% %v = 0\n", i, 2)
	}

	i = 0
	switch i {
	case 0:
		fmt.Println("Zero")
		fallthrough
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown")
	}
}

func argumentEvaluator(text string) string {
	return "Hello " + text
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

	fmt.Println("Now:", time.Now())
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
