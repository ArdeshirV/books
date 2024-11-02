package main

// book - The Go Programming Language

import (
	"fmt"
)

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

func main() {
	defer func() {
		//var tmp string
		fmt.Printf("%s\nPress enter to quit.\n", NORMAL)
		//fmt.Scanf("%s", &tmp)
	}()
	PrintTitle()
	fmt.Printf("    %sThe Go Programming Language %sʕ◔ϖ◔ʔ%s\n\n", BMAGENTA, BGREEN, TEAL)
	chapterOne()
}

func chapterOne() {
	fmt.Println("  Chapter One: ")
}

func PrintTitle() {
	blnColor := true
	strAppName := "Go.go"
	strAppYear := "2023"
	strAppDescription := "Test Golang codes"
	strVersion := "1.0"
	strLicense := "GPLv3+"
	strCopyright := "https://github.com/ArdeshirV"
	print(FormatTitle(strAppName, strAppDescription, strVersion, blnColor))
	print(FormatCopyright(strAppYear, strCopyright, strLicense, blnColor))
}
