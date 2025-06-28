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

func stepTwo() {
	printTitle("stepTwo")
	fmt.Println("The output of stepTwo also goes here.")
}

func stepOne() {
	printTitle("stepOne")

	const x, y = 1024, 2
	const TestNum = x<<8 + y<<16
	fmt.Printf("TestNum = %v\n", TestNum)

	var ce CustomEntity
	ce.SetOwner("Ardeshir")
	fmt.Printf("ce.Owner() = %s\n", ce.Owner())
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
