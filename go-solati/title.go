package main

import "fmt"

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
