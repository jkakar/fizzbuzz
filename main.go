package main

import "fmt"

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
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d: %s\n", i, fizzBuzz(i))
	}
}
