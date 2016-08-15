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
})
