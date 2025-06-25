package main

// book - Introduce Go

import (
	"container/list"
	"encoding/gob"
	"fmt"
	"hash/crc32"
	"io"
	"io/fs"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/ArdeshirV/books/introducing-go/mypackage"
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
	//Chapter6()
	//Chapter7()
	//Chapter8Packages()
	//Chapter8PackagesPart2()
	//Chapter8PackagesPart3HttpServer()
	//Chapter8PackagesPart4RPC()
	//Chapter8PackagesPart5()
	Chapter10Goroutines()
}

func Chapter10Goroutines() {
	//stepOne()
	//stepTwo()
	//stepThree()
	//stepFour()
	stepFive()
	//testMinimumFunc()
}

func testMinimumFunc() {
	list := []int{10, 20, 200, 2, 3, 9, 1, 12, 33, -1, 24, 34, -90, 20}
	fmt.Printf("Minimum(%v) = %v\n", list, Minimum(list))
}

func Minimum(list []int) int {
	var min int
	for i := 0; i < len(list); i++ {
		if i == 0 || min > list[i] {
			min = list[i]
		}
	}
	return min
}

func stepFive() {
	type HomePageSize struct {
		URL  string
		Size int
	}

	urls := [4]string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}
	fmt.Println(urls)

	results := make(chan HomePageSize)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize
	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}
	fmt.Println("The biggest home page:", biggest.URL)

	var input string
	fmt.Scanln(&input)
}

func stepFour() {
	fmt.Printf("\033[1;%dmChapter 10 - step 3\033[0;0m\n", 36)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg := <-c1:
				fmt.Println(msg)
			case msg := <-c2:
				fmt.Println(msg)
			case <-time.After(time.Second * 1):
				fmt.Println("timeout")
			default:
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

func stepThree() {
	for i := 0; i < 55; i++ {
		fmt.Printf("%d:\033[1;%dmChapter 10 - step 3\033[0;0m\n", i, i)
	}
}

func stepTwo() {
	var c chan string = make(chan string, 10)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func stepOne() {
	for i := 0; i < 10; i++ {
		go goF(i)
	}
	var input string
	fmt.Scanln(&input)
}

func goF(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func Chapter8PackagesPart5() {
	mypackage.Display("Hello from a mypackage!")
	max := mypackage.Max([]float64{10, 20, 40, 60, 90, 30})
	mypackage.Display(fmt.Sprintf("max:%f\n", max))
}

func Chapter8PackagesPart4RPC() {
	go server()
	go client()
	var input string
	fmt.Scanln(&input)
	os.Exit(0)
}

type Server struct{}

func (this *Server) Negate(name string, reply *string) error {
	*reply = "Hello dear " + name
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result string
	name := "Ardeshir"
	err = c.Call("Server.Negate", name, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(", name, ") =", result)
	}
}

func Chapter8PackagesPart3HttpServer() {
	http.HandleFunc("/golang", golangHandler)
	http.ListenAndServe(":9000", nil)
}

func golangHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>The Golang</title>
			</head>
			<body>
				The Golang Programming Language
			</body>
		</html>
		`,
	)
}

func Chapter8PackagesPart2() {
	go server()
	go client()
	var input string
	fmt.Scanln(&input)
}

func serverX() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received:", msg)
	}
	c.Close()
}

func clientX() {
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "The Golang Programming Language"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func Chapter8Packages() {
	fmt.Println("Chapter 8 - Packages")
	s1, s2 := "Ardeshir اردشیر", "shir"
	fmt.Println(s1, s2)
	fmt.Println(strings.EqualFold("Shir", "sHiR"))
	print(strings.Count(s1, "s"))
	fmt.Println(strings.Join([]string{s1, s2}, "x"))
	fmt.Println(strings.Repeat("Shir", 10))
	fmt.Println(strings.Fields("Ardeshir is   a computer    programmer.  "))
	fmt.Println(strings.IndexByte(s1, 66))
	fmt.Println(strings.ToTitle("Ardeshir Varmazyahr"))
	fmt.Println(strings.NewReader(s1))

	file, err := os.Open("/home/asohishn/.bashrc")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println(
		"stat.IsDir()", stat.IsDir(), "\n",
		"stat.Mode()", stat.Mode(), "\n",
		"stat.ModTime()", stat.ModTime(), "\n",
		"stat.Name()", stat.Name(), "\n",
		"stat.Size()", stat.Size(), "\n",
		"stat.Sys()", stat.Sys(), "\n")

	buff := make([]byte, stat.Size())
	_, err = file.Read(buff)
	if err != nil {
		panic(err)
	}

	bashrc := string(buff)
	fmt.Println(bashrc[0:200])

	of, err := os.Create("temp.txt")
	if err != nil {
		panic(err)
	}
	defer of.Close()
	of.WriteString(bashrc[:500])

	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	filesInDir, err := dir.ReadDir(-1)
	if err != nil {
		panic(err)
	}

	for _, fi := range filesInDir {
		fmt.Println(fi.Name())
	}

	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	profileFileName := os.Getenv("HOME") + "/.profile"
	profile, err := ReadFromFile(profileFileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(profileFileName)
		fmt.Println(profile)
	}

	fmt.Printf("toUpper(%s) = \"%s\"\n", s1, toUpper(s1))
	fmt.Printf("toLower(%s) = \"%s\"\n", s1, toLower(s1))

	if err = WriteToFile("/home/asohishn/Documents/Temp/Out.txt",
		"I love you!"); err != nil {
		fmt.Println(err)
	}

	listTest := list.New()
	listTest.PushBack(10)
	listTest.PushBack(20)
	listTest.PushBack(30)
	fmt.Println(listTest)
	for element := listTest.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	names := []Name{
		{FirstName: "Mohammad", LastName: "Aghaee"},
		{FirstName: "Ardeshir", LastName: "Varmazyahr"},
	}
	data := Names(names)
	fmt.Println(data)
	sort.Sort(data)
	fmt.Println(data)

	name := Name{FirstName: "Abo", LastName: "Sofhiyahn"}
	fmt.Println(name)
	pname := &name
	fmt.Println(pname)

	h := crc32.NewIEEE()
	strName := fmt.Sprint(data[0])
	fmt.Println(strName)
	h.Write([]byte(strName))
	v32 := h.Sum32()
	fmt.Println(v32)

	//sha := crypto.Hash.New(crypto.MD4)
	//sha.Write([]byte("Ardeshir Varmazyahr"))
	//b := sha.Sum([]byte{})
	//fmt.Println(b)
}

type Name struct {
	FirstName, LastName string
}

func (this Name) String() string {
	return fmt.Sprintf("[%s, %s]", this.FirstName, this.LastName)
}

type Names []Name

func (this Names) Len() int {
	return len(this)
}

func (this Names) Less(first, second int) bool {
	return this[first].FirstName < this[second].FirstName
}

func (this Names) Swap(left, right int) {
	this[left].FirstName, this[right].FirstName =
		this[right].FirstName, this[left].FirstName
	this[left].LastName, this[right].LastName =
		this[right].LastName, this[left].LastName
}

func WriteToFile(fileName, contents string) error {
	fileHandle, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	if _, err = fileHandle.WriteString(contents); err != nil {
		return err
	}

	return nil
}

func toUpper(input string) string {
	var b strings.Builder
	const distance = 'a' - 'A'
	for _, ch := range input {
		if ch >= 'a' && ch <= 'z' {
			b.WriteRune(ch - distance)
		} else {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func toLower(input string) string {
	var b strings.Builder
	const distance = 'a' - 'A'
	for _, ch := range input {
		if ch >= 'A' && ch <= 'Z' {
			b.WriteRune(ch | distance)
		} else {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func ReadFromFile(fileName string) (contents string, err error) {
	fileHandle, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()

	stat, err := fileHandle.Stat()
	if err != nil {
		return "", err
	}

	buff := make([]byte, stat.Size())
	_, err = fileHandle.Read(buff)
	if err != nil {
		return "", err
	}
	contents = string(buff)

	return contents, nil
}

func print(args ...any) {
	fmt.Println(args...)
}

func Chapter7() {
	var s Shape
	fmt.Printf("Type of s: %T\n", s)
	s = &Rectangel{TopLeft: Point{10, 10}, BottomRight: Point{30, 30}}
	fmt.Println(s.Area())
	var sq Square
	s = &sq
	fmt.Println(s, s.Area())
	fmt.Println(sq, sq.Area())
	fmt.Printf("Type of s: %T\n", s)
}

type Square struct {
	Rectangel
}

func (this *Square) Area() float64 {
	return 0
}

func (this *Square) Perimeter() float64 {
	return 4321.0
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Point struct {
	X, Y float64
}

type Rectangel struct {
	TopLeft, BottomRight Point
}

func (s *Rectangel) Area() float64 {
	return 100.0
}

func (this *Rectangel) Perimeter() float64 {
	return 1234
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

	arrNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := Sum(arrNums...)
	fmt.Println(sum)

	v, b := Half(1)
	fmt.Println(v, b)
	v, b = Half(2)
	fmt.Println(v, b)

	fmt.Println(FindMax(arrNums))

	odd := makeOddGenerator()
	for i := 0; i < 10; i++ {
		fmt.Print(odd(), " ")
	}
	fmt.Println()

	for i := 1; i < 10; i++ {
		fmt.Print(fib(i), " ")
	}
	fmt.Println()

	a1, b1 := 10, 20
	fmt.Println(a1, b1)
	swap(&a1, &b1)
	fmt.Println(a1, b1)
	a1, b1 = b1, a1
	fmt.Println(a1, b1)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func makeOddGenerator() func() int {
	number := 1
	return func() int {
		number += 2
		return number
	}
}

func FindMax(list []int) int {
	result := 0
	for _, value := range list {
		if value > result {
			result = value
		}
	}
	return result
}

func Half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func Sum(arr ...int) (total int) {
	total = 0
	for _, value := range arr {
		total += value
	}
	return total
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

