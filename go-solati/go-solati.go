// go-solati.go: My practices about "The Go programming language reference" by Mosfata Solati
package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Printf("%s\nPress enter to quit.", NORMAL)

		var tmp string
		fmt.Scanf("%s", &tmp)
	}()
	PrintTitle()
	bookName := "The Go programming language reference"
	title := "\n    %sMy Practices about \"%s%s%s\" by Mostafa Solati %sʕ◔ϖ◔ʔ%s\n\n"
	fmt.Printf(title, MAGENTA, BMAGENTA, bookName, MAGENTA, BGREEN, TEAL)

	mainChapterOne()
}

func mainChapterOne() {
	fmt.Println("  Chapter One: ")
}
