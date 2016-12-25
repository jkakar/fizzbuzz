package main

import (
	"fmt"
	"testing"
)

// Values evenly divisible by 3 produce a Fizz tesult.
func TestFizz(t *testing.T) {
	for _, n := range []int{3, 6, 9} {
		r := fizzBuzz(n)
		if r != "Fizz" {
			t.Fatalf("Expected Fizz for %d but got %s", n, r)
		}
	}
}

// Values evenly divisible by 5 produce a Buzz tesult.
func TestBuzz(t *testing.T) {
	for _, n := range []int{5, 10, 20} {
		r := fizzBuzz(n)
		if r != "Buzz" {
			t.Fatalf("Expected Buzz for %d but got %s", n, r)
		}
	}
}

// Values evenly divisible by both 3 and 5 produce a FizzBuzz tesult.
func TestFizzBuzz(t *testing.T) {
	for _, n := range []int{15, 30, 45} {
		r := fizzBuzz(n)
		if r != "FizzBuzz" {
			t.Fatalf("Expected FizzBuzz for %d but got %s", n, r)
		}
	}
}

// Values not divisble by 3, 5 or by 3 and 5 simultaneously return the numeric
// value.
func TestOther(t *testing.T) {
	for _, n := range []int{4, 8, 13} {
		r := fizzBuzz(n)
		if r != fmt.Sprintf("%d", n) {
			t.Fatalf("Expected %d for %d but got %s", n, n, r)
		}
	}
}
