package main

import (
	"fmt"
)

func main() {
	var i int
	for {
		i++
		switch {
		case i%15 == 0:
			fmt.Println(i, ":FizzBuzz")
		case i%3 == 0:
			fmt.Println(i, ":Fizz")
		case i%5 == 0:
			fmt.Println(i, ":Buzz")
		}
		if i == 100 {
			break
		}
	}
}
