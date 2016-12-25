package main

import (
	"flag"
	"fmt"
)

var iterations = flag.Int("i", 100, "number of iterations to calculate")

func fizzBuzz(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}

func main() {
	flag.Parse()
	for i := 1; i <= *iterations; i++ {
		fmt.Printf("%d: %s\n", i, fizzBuzz(i))
	}
}
