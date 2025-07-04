package main

import (
	"fmt"
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

	stepTwo()
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

func (ce *CustomEntity) Owner() string {
	return ce.owner
}

func (ce *CustomEntity) SetOwner(owner string) {
	ce.owner = owner
}

func (ce CustomEntity) String() string {
	return ce.owner
}
