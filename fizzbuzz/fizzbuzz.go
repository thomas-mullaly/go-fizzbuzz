package fizzbuzz

import "fmt"

// Range used by FizzBuzz program
type Range struct {
	Start, End int
}

type FizzBuzz struct {
	rules []Rule
}

// NewFizzBuzz creates a new FizzBuzz object
func NewFizzBuzz() *FizzBuzz {
	return &FizzBuzz{
		rules: []Rule{
			NewDoubleModuloRule(3, 5, "FizzBuzz"),
			NewSingleModuloRule(3, "Fizz"),
			NewSingleModuloRule(5, "Buzz"),
			&NumberEchoRule{},
		},
	}
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
		for _, rule := range fb.rules {
			if rule.Applies(i) {
				value = rule.Result(i)
				break
			}
		}

		output = append(output, value)
	}

	return output, nil
}
