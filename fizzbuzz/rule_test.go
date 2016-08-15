package fizzbuzz

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRule(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rules Suite")
}

var _ = Describe("SingleModuloRule", func() {
	var (
		rule *SingleModuloRule
	)

	BeforeEach(func() {
		rule = NewSingleModuloRule(3, "Fizz")
	})

	Describe("Applies", func() {
		It("Should return false if value is not divisble", func() {
			Expect(rule.Applies(5)).To(Equal(false))
		})

		It("Should return true if value is divisible", func() {
			Expect(rule.Applies(6)).To(Equal(true))
		})
	})

	Describe("Result", func() {
		It("Should return 'Fizz'", func() {
			Expect(rule.Result(1)).To(Equal("Fizz"))
		})
	})
})

var _ = Describe("DoubleModuloRule", func() {
	var (
		rule *DoubleModuloRule
	)

	BeforeEach(func() {
		rule = NewDoubleModuloRule(3, 5, "FizzBuzz")
	})

	Describe("Applies", func() {
		It("Should return false if number is not divisible by three and five", func() {
			Expect(rule.Applies(12)).To(Equal(false))
		})

		It("Should return true if number is divisible by three and five", func() {
			Expect(rule.Applies(15)).To(Equal(true))
		})
	})

	Describe("Result", func() {
		It("Should return FizzBuzz", func() {
			Expect(rule.Result(15)).To(Equal("FizzBuzz"))
		})
	})
})

var _ = Describe("NumberEchoRule", func() {
	var (
		rule *NumberEchoRule
	)

	BeforeEach(func() {
		rule = &NumberEchoRule{}
	})

	Describe("Applies", func() {
		It("Should always return true", func() {
			Expect(rule.Applies(1)).To(Equal(true))
			Expect(rule.Applies(3)).To(Equal(true))
		})
	})

	Describe("Result", func() {
		It("Should return string form of input", func() {
			Expect(rule.Result(13)).To(Equal("13"))
			Expect(rule.Result(2)).To(Equal("2"))
		})
	})
})
