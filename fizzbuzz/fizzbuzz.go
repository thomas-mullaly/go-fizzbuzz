package fizzbuzz

import (
	"fmt"
	"strconv"
)

// Range used by FizzBuzz program
type Range struct {
	Start, End int
}

type FizzBuzz struct{}

// NewFizzBuzz creates a new FizzBuzz object
func NewFizzBuzz() *FizzBuzz {
	return &FizzBuzz{}
}

// Run runs the fizzbuzz logic on the given range
func (fb *FizzBuzz) Run(r Range) ([]string, error) {
	var output []string

	if r.Start > r.End {
		return nil, fmt.Errorf("Start (%d) cannot be greater than End (%d)", r.Start, r.End)
	}

	output = make([]string, 0, r.End-r.Start)

	for i := r.Start; i < r.End; i++ {
		value := ""
		if i%3 == 0 && i%5 == 0 {
			value = "FizzBuzz"
		} else if i%3 == 0 {
			value = "Fizz"
		} else if i%5 == 0 {
			value = "Buzz"
		} else {
			value = strconv.Itoa(i)
		}

		output = append(output, value)
	}

	return output, nil
}
