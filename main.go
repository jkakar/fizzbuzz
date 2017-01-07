package main

import (
	"flag"
	"fmt"
	"strconv"
)

var iterations = flag.Int("i", 100, "number of iterations to calculate")

func fizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

func main() {
	flag.Parse()
	for i := 1; i <= *iterations; i++ {
		fmt.Printf("%d: %s\n", i, fizzBuzz(i))
	}
}
