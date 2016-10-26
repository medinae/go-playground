// Fizz Buzz with if-else

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {

		if i%15 == 0 {
			fmt.Println("Fizz-Buzz", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz", i)
		} else if i%3 == 0 {
			fmt.Println("fizz", i)
		} else {
			fmt.Println(i)
		}
	}
}
