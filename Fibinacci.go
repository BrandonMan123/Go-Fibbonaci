package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func(int) int {

	a := 0
	b := 1
	c := 0

	return func(x int) int {

		c = a + b
		a = b
		b = c
		return c

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(2))
	}
}
