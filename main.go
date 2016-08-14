package main

import "github.com/thomas-mullaly/go-fizzbuzz/fizzbuzz"

func main() {
	fizzBuzz := fizzbuzz.NewFizzBuzz()

	fizzBuzz.Run(fizzbuzz.Range{Start: 1, End: 4})
}
