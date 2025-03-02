// go-solati.go: My practices about "The Go programming language reference"
package main

import (
	"fmt"
)

func main() {
	defer func() {
		//waiteForEnter()
	}()
	PrintTitle()
	bookName := "The Go programming language reference"
	title := "\n    %sMy Practices about \"%s%s%s\" %sʕ◔ϖ◔ʔ%s\n\n"
	fmt.Printf(title, MAGENTA, BMAGENTA, bookName, MAGENTA, BGREEN, TEAL)

	mainChapterOne()
	//mainChapterTwo()
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
	fmt.Println("\003[0;35mاردشیر")
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
//
