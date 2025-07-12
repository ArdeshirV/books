package main

import (
	"fmt"
	"os"
)

func printTitle(title string) {
	fmt.Printf("\n  \033[1;36m%s\033[0;36m\n", title)
}

func printMainTitle() {
	fmt.Print("\n    \033[1;35mMy practices about ",
		"\033[0;32m\"Effective Go\" \033[1;32mʕ◔ϖ◔ʔ\033[0;36m\n")
}

func main() {
	defer func() {
		fmt.Print("\033[0m")
	}()
	printMainTitle()
	//stepOne()
	//stepTwo()
	//stepThree()
	//stepFour()
	//stepFive()
	stepSix()
}

func stepSix() {
	printTitle("stepSix\n\n")

	fmt.Printf("Hello, World!\n\n")
}

func stepFive() {
	printTitle("stepFive\n\n")

	useSemaphore()
}

// A buffered channel can be used like a semaphore, for instance to limit throughput.
func useSemaphore() {
	request := &Request{args: []int{3, 4, 5}, f: sum, resultChan: make(chan int)}
	clientRequests := make(chan *Request, MaxOutstanding)
	clientRequests <- request
	isFinished := make(chan bool)
	Serve(&clientRequests, isFinished)
	<-isFinished
	fmt.Printf("Answer: %d\n", <-request.resultChan)
}

const MaxOutstanding = 20

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func handle(r *Request) {
	r.resultChan <- r.f(r.args)
}

func Serve(clientRequests *chan *Request, quit chan bool) {
	for req := range *clientRequests {
		go handle(req)
	}
	<-quit
}

func stepFour() {
	printTitle("stepFour\n")

	fileName := "/home/asohishn/Documents/Projects/Go/books/effective-go/effective-go.go"
	if fileInfo, err := os.Stat(fileName); err == nil {
		fmt.Printf("fileInfo: %#v\n\n", fileInfo)
	}

	c := make(chan int)
	go func() {
		fmt.Println("First - in Goroutine")
		for i := range 20 {
			fmt.Print(i, " ")
		}
		fmt.Println()
		c <- 1
	}()
	fmt.Println("Second - idle time")
	<-c
	fmt.Printf("Third - in Main thread\n\n")
}

func stepThree() {
	printTitle("stepThree\n")
	someSlice := []string{"Hi", "Hallo", "Guten Tag"}
	fmt.Println(someSlice)
	anotherSlice := [2][3]string{{"One", "Two", "Three"}, {"Yek", "Do", "Se"}}
	fmt.Println(anotherSlice)

	m1 := make(map[string]int, 10)
	m1["Hello"] = 10
	fmt.Println(m1)
	m2 := make(map[string]int, 0)
	m2["Hi"] = 100
	fmt.Println(m2, ", m1[\"Guten Morgen\"] = ", m1["Guten Morgen"])

	fmt.Println("User =", os.Getenv("USER"))

	var c CustomEntity
	c.SetOwner("Ardeshir")
	fmt.Fprint(&c, " Ki Chihr Hach Yazataan")
	fmt.Printf("c: %v\n\n", c)

	IsString("Hello, World!")
	IsString(190.003)
	fmt.Println()

}

func IsString(value any) {
	if str, ok := value.(string); ok {
		fmt.Printf("string value is: %q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}

func init() {
	fmt.Println("effective-go.init() ...")
}

func stepTwo() {
	printTitle("stepTwo\n")

	fmt.Printf("%c, %c\n", 'A', 65)
	fmt.Println("string(65) =", string(rune(65)))
	fmt.Println("rune(65) =", rune(65))
	fmt.Printf("byte(65) = %v\n\n", byte(65))

	var c CustomEntity
	c.SetOwner("Someone")
	fmt.Printf("%v\n\n", c)

	c1 := new(CustomEntity)
	c1.SetOwner("user")
	c2 := &CustomEntity{owner: "user"}
	fmt.Printf("c1 = %v\nc2 = %v\n\n", c1, c2)

	const (
		Enone  = 5
		Eio    = 1
		Einval = 2
	)
	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	fmt.Printf("T(a) = %T, a = %v\nT(s) = %T, s = %v\nT(m) = %T, m = %v\n\n",
		a, a, s, s, m, m)

	arr1 := []int{10, 20, 30}
	arr2 := arr1
	arr2[0] = 1000
	arr1[1] = 100
	fmt.Printf("arr1 = %v, arr2 = %v\n\n", arr1, arr2)

	fmt.Printf("c: %v\n", c)
	fmt.Fprint(&c, " Hello")
	fmt.Printf("c: %v\n", c)
}

func stepOne() {
	printTitle("stepOne\n")

	const x, y = 1024, 2
	const TestNum = x<<8 + y<<16
	fmt.Printf("TestNum = %v\n\n", TestNum)

	var ce CustomEntity
	ce.SetOwner("Ardeshir")
	fmt.Printf("ce.Owner() = %s\n\n", ce.Owner())

	someString := "اردشیر"
	for i, v := range someString {
		fmt.Printf("str[%d] = %#U, %v\n", i, v, v)
	}
	for i := 0; i < len(someString); i++ {
		fmt.Printf("str[%d] = %#U, %v\n",
			i, someString[i], someString[i])
	}
	fmt.Println()

	m, n := 10, 20
	fmt.Printf("m = %v, n = %v\n", m, n)
	i, n := 100, 200
	fmt.Printf("i = %v, n = %v\n\n", i, n)

	fmt.Printf("'%v'\n", someString)
	reversedSomeString := []byte(someString)
	for i, j := 0, len(reversedSomeString)-1; i < j; i, j = i+1, j-1 {
		reversedSomeString[i] = reversedSomeString[j]
	}
	fmt.Printf("'%v'\n", string(reversedSomeString))
	reversed2 := []rune(someString)
	for i, j := 0, len(reversed2)-1; i < j; i, j = i+1, j-1 {
		reversed2[i] = reversed2[j]
	}
	fmt.Printf("'%v'\n", string(reversed2))
	reversed3 := []rune(string(reversed2))
	for i, j := 0, len(reversed3)-1; i < j; i, j = i+1, j-1 {
		reversed3[i] = reversed3[j]
	}
	fmt.Printf("'%v'\n\n", string(reversed3))

	fmt.Printf("ValueOf('A') = %d\n\n", 'A')

	hexString := "A20D1"
	fmt.Printf("Decimal('%v') = %v\n", hexString, hex2Decimal(hexString))
	fmt.Printf("Decimal2('%v') = %v\n\n", hexString, hex2Decimal2(hexString))

	// Type Switch
	t := any(unhex)
	// The 'v' has different type in each case
	switch v := t.(type) {
	case bool:
		fmt.Printf("bool\n\n")
	default:
		fmt.Printf("type(t) = %T\n\n", v)
	}
}

func hex2Decimal2(hexString string) (number int) {
	for _, c := range []byte(hexString) {
		number *= 16
		number += int(unhex(c))
	}
	return
}

func hex2Decimal(hexString string) int {
	index, result := 1, 0
	reversedHexString := []byte(reverse(hexString))
	for _, c := range reversedHexString {
		result += int(unhex(c)) * index
		index *= 16
	}
	return result
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

type CustomEntity struct {
	owner string
}

func (ce *CustomEntity) Write(data []byte) (int, error) {
	ce.owner += string(data)
	return len(data), nil
}

func (ce CustomEntity) Owner() string {
	return ce.owner
}

func (ce *CustomEntity) SetOwner(owner string) {
	ce.owner = owner
}

func (ce CustomEntity) String() string {
	return ce.owner
}

