package fizzbuzz

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFizzBuzz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rules Suite")
}

func verifyFizzBuzzResults(actual []string, expected []string) {
	expectedLen := len(expected)
	actualLen := len(actual)

	Expect(actualLen).To(Equal(expectedLen))

	for index, value := range expected {
		Expect(actual[index]).To(Equal(value))
	}
}

var _ = Describe("FizzBuzz", func() {
	var (
		fizzBuzz *FizzBuzz
	)

	BeforeEach(func() {
		fizzBuzz = NewFizzBuzz()
	})

	Describe("Run", func() {
		It("Should return Fizz for numbers divisible by three", func() {
			numRange := Range{Start: 3, End: 4}

			result, _ := fizzBuzz.Run(numRange)

			verifyFizzBuzzResults(result, []string{"Fizz"})
		})

		It("Should return actual number when not divisible by three or five", func() {
			numRange := Range{Start: 2, End: 3}

			result, _ := fizzBuzz.Run(numRange)

			verifyFizzBuzzResults(result, []string{"2"})
		})

		It("Should return Buzz when number is divisible by five", func() {
			numRange := Range{Start: 5, End: 6}

			result, _ := fizzBuzz.Run(numRange)

			verifyFizzBuzzResults(result, []string{"Buzz"})
		})

		It("Should return FizzBuzz when number is divisible by three and five", func() {
			numRange := Range{Start: 15, End: 16}

			result, _ := fizzBuzz.Run(numRange)

			verifyFizzBuzzResults(result, []string{"FizzBuzz"})
		})

		It("Should return error when Start is greater than End", func() {
			numRange := Range{Start: 2, End: 1}

			result, err := fizzBuzz.Run(numRange)

			Expect(result).To(BeEmpty())
			Expect(err).ToNot(Equal(nil))
		})
	})
})
