package main
// book - The Go Programming Language
//01234567890

import (
	"fmt"
)

func main() {
	defer func() {
		//var tmp string
		fmt.Printf("%s\nPress enter to quit.\n", NORMAL)
		//fmt.Scanf("%s", &tmp)
	}()
	PerformTitle()
	chapter1()
}

func chapter1() {
	printTitle("  Chapter 1:\n")
	//
}

func chapter2() {
	printTitle("  Chapter 2:\n")
}

func chapter3() {
	printTitle("  Chapter 3:\n")
}

func chapter4() {
	printTitle("  Chapter 4:\n")
}

func chapter5() {
	printTitle("  Chapter 5:\n")
}

func chapter6() {
	printTitle("  Chapter 6:\n")
}

func chapter7() {
	printTitle("  Chapter 7:\n")
}

func chapter8() {
	printTitle("  Chapter 8:\n")
}

func chapter9() {
	printTitle("  Chapter 9:\n")
}

func chapter10() {
	printTitle("  Chapter 10:\n")
}

func chapter11() {
	printTitle("  Chapter 11:\n")
}

func chapter12() {
	printTitle("  Chapter 12:\n")
}

func chapter13() {
	printTitle("  Chapter 13:\n")
}

func chapter14() {
	printTitle("  Chapter 14:\n")
}

func chapter15() {
	printTitle("  Chapter 15:\n")
}

func PrintTitle() {
	blnColor := true
	strAppName := "the-go-programming-language"
	strAppYear := "2024"
	strAppDescription := "The Go Programming Language"
	strVersion := "1.0"
	strLicense := "GPLv3+"
	strCopyright := "https://github.com/ArdeshirV"
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
		"\033[0;33m, Licensed under \033[1;33m%v\033[0m\n")
	var strFormat string
	if blnColor {
		strFormat = Colored
	} else {
		strFormat = NoneColored
	}
	return sprintf(strFormat, strAppYear, strCopyright, strLicense)
}

func sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func PerformTitle() {
	PrintTitle()
	title := "\n    %sThe Go Programming Language %sʕ◔ϖ◔ʔ%s\n\n"
	fmt.Printf(title, BMAGENTA, BGREEN, TEAL)
}

func printTitle(title string) {
	fmt.Printf("\n  \033[1;36m%s\033[0;36m\n", title)
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

