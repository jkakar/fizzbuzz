package main

import (
	"flag"
	"fmt"
	"strconv"
)

func FizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", n)
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		panic("expected a single numeric argument")
	}
	n, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	fmt.Println(FizzBuzz(n))
}
