package main

// book - Effective Go

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
		fmt.Printf("%s\nPress enter to quit.\n", NORMAL)
		//var tmp string
		//fmt.Scanf("%s", &tmp)
	}()
	PrintTitle()
	const title = "\n    %sEffective Go %sʕ◔ϖ◔ʔ%s\n\n"
	fmt.Printf(title, BMAGENTA, BGREEN, TEAL)
	chapterOne()
}

func chapterOne() {
	fmt.Println("  Chapter One: ")
}

func PrintTitle() {
	blnColor := true
	strAppName := "effective-go"
	strAppYear := "2025"
	strAppDescription := "Effective Go"
	strVersion := "1.0"
	strLicense := "GPLv3+"
	strCopyright := "https://github.com/ArdeshirV"
	fmt.Print(FormatTitle(strAppName, strAppDescription, strVersion, blnColor))
	fmt.Print(FormatCopyright(strAppYear, strCopyright, strLicense, blnColor))
}

func FormatTitle(strAppName, strAppDescription, strVersion string,
	blnColor bool) string {
	NoneColored := "%v - %v Version %v\n"
	Colored := "\033[1;33m%v\033[0;33m - %v \033[1;33mVersion %v\033[0m\n"
	var strFormat string
	if blnColor {
		strFormat = Colored
	} else {
		strFormat = NoneColored
	}
	return fmt.Sprintf(strFormat, strAppName, strAppDescription, strVersion)
}

func FormatCopyright(strAppYear, strCopyright, strLicense string,
	blnColor bool) string {
	NoneColored := "Copyright (c) %v %v, Licensed under %v\n"
	Colored := ("\033[0;33mCopyright (c) \033[1;33m%v \033[1;34m%v" +
		"\033[0;33m, Licensed under \033[1;33m%v\033[0m\n")
	var strFormat string
	if blnColor {
		strFormat = Colored
	} else {
		strFormat = NoneColored
	}
	return fmt.Sprintf(strFormat, strAppYear, strCopyright, strLicense)
}
