package fizzbuzz

import "testing"

func verifyFizzBuzzResults(actual []string, expected []string, t *testing.T) {
	expectedLen := len(expected)
	actualLen := len(actual)

	if actualLen != expectedLen {
		t.Fatalf("Expected %d results but got %d", expectedLen, actualLen)
	}

	for index, value := range expected {
		if actual[index] != value {
			t.Fatalf("Expected '%s' at index %d but got '%s'", value, index, actual[index])
		}
	}
}

func TestReturnFizzForNumbersDivisbleByThree(t *testing.T) {
	fizz := NewFizzBuzz()

	numRange := Range{Start: 3, End: 4}

	result, err := fizz.Run(numRange)

	if err != nil {
		t.Fatalf("Expected err to be nil but was %s instead", err)
	}

	verifyFizzBuzzResults(result, []string{"Fizz"}, t)
}

func TestReturnActualNumberWhenNotDivisbleByThree(t *testing.T) {
	fizz := NewFizzBuzz()

	numRange := Range{Start: 2, End: 3}

	result, err := fizz.Run(numRange)

	if err != nil {
		t.Fatalf("Expected err to be nil but was %s instead", err)
	}

	verifyFizzBuzzResults(result, []string{"2"}, t)
}

func TestReturnBuzzWhenNumberIsDivisbleByFive(t *testing.T) {
	fizz := NewFizzBuzz()

	numRange := Range{Start: 5, End: 6}

	result, err := fizz.Run(numRange)

	if err != nil {
		t.Fatalf("Expected err to be nil but was %s instead", err)
	}

	verifyFizzBuzzResults(result, []string{"Buzz"}, t)
}

func TestReturnErrIfStartIsGreaterThanEnd(t *testing.T) {
	fizz := NewFizzBuzz()

	numRange := Range{Start: 2, End: 1}

	result, err := fizz.Run(numRange)

	if result != nil {
		t.Fatal("Expected result to be nil")
	}

	if err == nil {
		t.Fatalf("Expected err to not be nil")
	}
}

func TestReturnFizzBuzzIfNumberIsDivisibleByThreeAndFive(t *testing.T) {
	fizz := NewFizzBuzz()

	numRange := Range{Start: 15, End: 16}

	result, _ := fizz.Run(numRange)

	verifyFizzBuzzResults(result, []string{"FizzBuzz"}, t)
}
