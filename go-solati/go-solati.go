// go-solati.go: My practices about "The Go programming language reference by Solati"
package main

import (
	"fmt"
	"sort"
)

// Main entry point
func main() {
	defer func() {
		//waiteForEnter()
	}()
	PerformTitle()

	//mainChapterOne()
	//mainChapterTwo()
	mainChapterThree()
}

// Chapter three
func mainChapterThree() {
	fmt.Printf("  Chapter Three: \n\n")
	fmt.Println("Hello World")
	Integrate("Hello from mainChapterThree +/-")
	// TODO: add your new code here
}

// Integrate ax^2+bx+c
func Integrate(expr string) float64 {
	fmt.Println(expr)
	return 0.0
}

// find n max of array
func findNMax(arr []int, n int) []int {
	if n > len(arr) {
		n = len(arr)
	}
	result := make([]int, n)
	copy(result, arr)
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result[:n]
}

func findMaxPrime(n int) int {
	maxPrime := 2
	for i := 3; i <= n; i++ {
		if isPrime(i) && i > maxPrime {
			maxPrime = i
		}
	}
	return maxPrime
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func rev(text string) string {
	return text
}

func EvalMathematicExpr(expr string) float64 {
	fmt.Print("Hello")
	return 0.0
}

func integrate() {
	fmt.Println("&&&")
}

func mainChapterOne() {
	fmt.Printf("  Chapter One: \n\n")

	fmt.Println("Hello World")
	fmt.Println(`Hello, World!
in multilines.`)

	var age int
	var name string
	var movie string
	var score float64
	name = "Parmis"
	age = 13
	movie = "Hotel Transylvania"
	score = 7
	fmt.Println(name, "is a good student")
	fmt.Println(name, "is", age, "years old")
	fmt.Println(movie, name, "favorite movie score is", score)

	age = 39
	name = "اردشیر"
	fmt.Println(name, "برنامه نویس کامپیوتر است و ", age, "سال سن دارد")

	const (
		a = iota + 1
		b
		c
		_
		e
	)
	fmt.Println(a, b, c, e)

	var bit1 uint8 = 0b11101101
	fmt.Printf("%b, %b\n", bit1, ^bit1)

	str1 := "Hello"
	str2 := "World"
	str3 := fmt.Sprintf("%s, %s!", str1, str2)
	fmt.Println(str3)

	// TODO: Add your code here://
	fmt.Printf("%v\n", str1 != str2)
	const bit = 24
	fmt.Println(bit >> 2)
	const b1, b2 = 0b00110000, 0b00100001
	fmt.Printf("%b\n", b1^b2)

	gender := "male"
	if gender == "female" {
		fmt.Println("You are a woman")
	} else {
		fmt.Println("You are a man")
	}

	const day = 2
	dayName := ""

	switch day {
	case 0:
		dayName = "Saturday"
	case 1:
		dayName = "Sunday"
	case 2:
		dayName = "Monday"
	case 3:
		dayName = "Tuesday"
	case 4:
		dayName = "Wednsday"
	case 5:
		dayName = "Tursday"
	case 6:
		dayName = "Friday"
	default:
		dayName = "Xday"
	}

	arrDayNames := [7]string{"Saturday", "Sunday",
		"Monday", "Tuesday", "Wednsday", "Friday"}

	fmt.Println(dayName)
	fmt.Println(arrDayNames[day])

	goto next_statement
	fmt.Println("Khameneee is a dirty pig")
next_statement:
	fmt.Println("Yeah we are here")

	// Functions begins here:
	helloFunc()
	c1 := newCounter()
	c2 := newCounter()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c2())

	expr := "2*3+4/2*2-1"
	operators := make(map[string]func(int, int) int)
	operators["+"] = add
	operators["-"] = sub
	operators["/"] = div
	operators["*"] = mul
	res := 0
	prev := 0
	newExpr := ""
	for i, v := range expr {
		item := string(v)
		if item == "*" || item == "/" {
			next := int(expr[i+1])
			res = operators[item](prev, next)
			newExpr += fmt.Sprintf("%d ", res)
		}
		prev = int(v)
	}
	for i, v := range expr {
		item := string(v)
		if item == "+" || item == "-" {
			next := int(expr[i+1])
			res = operators[item](prev, next)
			newExpr += fmt.Sprintf("%d ", res)
		}
		prev = int(v)
	}
	fmt.Println("\n", newExpr)

	// Stone 3
	fmt.Println("Stone 3")

}

func mul(a, b int) int {
	return a * b
}

func div(b, a int) int {
	return b / a
}

func add(a, b int) int {
	return a + b
}

func sub(b, a int) int {
	return b - a
}

func newCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func helloFunc() {
	fmt.Println("helloFunc")
}

func mainChapterTwo() {
	fmt.Printf("  Chapter Two: \n\n")
	fmt.Println("Hello World")
}

func PrintTitle() {
	blnColor := true
	strAppName := "go-solati"
	strAppYear := "2025"
	strAppDescription := "The Go Programming Language"
	strVersion := "1.0"
	strLicense := "GPLv3+"
	strCopyright := "https://github.com/ArdeshirV/book/go-solati"
	fmt.Print(FormatTitle(strAppName, strAppDescription, strVersion, blnColor))
	fmt.Print(FormatCopyright(strAppYear, strCopyright, strLicense, blnColor))
}

func FormatTitle(strAppName, strAppDescription, strVersion string, blnColor bool) string {
	NoneColored := "%v - %v Version %v\n"
	Colored := "\033[1;33m%v\033[0;33m - %v \033[1;33mVersion %v\033[0m\n"
	var strFormat string
	if blnColor {
		strFormat = Colored
	} else {
		strFormat = NoneColored
	}
	return sprintf(strFormat, strAppName, strAppDescription, strVersion)
}

func FormatCopyright(strAppYear, strCopyright, strLicense string, blnColor bool) string {
	NoneColored := "Copyright (c) %v %v, Licensed under %v\n"
	Colored := ("\033[0;33mCopyright (c) \033[1;33m%v \033[1;34m%v" +
		"\033[0;33m, \033[1;33m%v\033[0m\n")
	var strFormat string
	if blnColor {
		strFormat = Colored
	} else {
		strFormat = NoneColored
	}
	return sprintf(strFormat, strAppYear, strCopyright, strLicense)
}

func sprintf(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

func PerformTitle() {
	PrintTitle()
	bookName := "The Go programming language reference"
	title := "\n    %sMy Practices about \"%s%s%s\" %sʕ◔ϖ◔ʔ%s\n\n"
	fmt.Printf(title, MAGENTA, BMAGENTA, bookName, MAGENTA, BGREEN, TEAL)
}

const (
	NORMAL   = "\033[0m"
	BOLD     = "\033[1m"
	RED      = "\033[0;31m"
	TEAL     = "\033[0;36m"
	WHITE    = "\033[0;37m"
	BLUE     = "\033[0;34m"
	GREEN    = "\033[0;32m"
	YELLOW   = "\033[0;33m"
	MAGENTA  = "\033[0;35m"
	BRED     = "\033[1;31m"
	BBLUE    = "\033[1;34m"
	BTEAL    = "\033[1;36m"
	BWHITE   = "\033[1;37m"
	BGREEN   = "\033[1;32m"
	BYELLOW  = "\033[1;33m"
	BMAGENTA = "\033[1;35m"
)

// Problems in book
// P29: sudo -C /usr/local -xzf flle-name --> tar -xvf file-name
// P44: پیدا کردن کوحکترین عدد برای یافتن بزرکترین عدد
// P55: true || false ==> false
// P54: a < b, a < b
// P79: تعریف تابغ به خط فارسی برعکس شده است
