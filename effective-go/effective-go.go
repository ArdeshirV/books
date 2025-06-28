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

	stepOne()
	//stepTwo()
}

func stepOne() {
	printTitle("stepOne")
	fmt.Println("The output goes here.")

}

func stepTwo() {
	printTitle("stepTwo")
	fmt.Println("The output of stepTwo also goes here.")
}
