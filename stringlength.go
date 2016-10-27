package main

import (
	"fmt"
	"strings"
)

func main() {
	stringOne := "merpflakes"
	stringTwo := "   "
	stringThree := ""

	if len(strings.TrimSpace(stringOne)) == 0 {
		fmt.Println("String is empty!")
	}

	if len(strings.TrimSpace(stringTwo)) == 0 {
		fmt.Println("String two is empty!")
	}

	if len(stringTwo) == 0 {
		fmt.Println("String two is still empty!")
	}

	if len(strings.TrimSpace(stringThree)) == 0 {
		fmt.Println("String three is empty!")
	}
}
