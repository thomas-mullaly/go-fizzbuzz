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
			Expect(rule.Applies(6)).To(Equal(true))
		})
	})
})

func TestSingleModuloRuleAppliesReturnsTrueIfDivisible(t *testing.T) {
	rule := NewSingleModuloRule(3, "Fizz")

	if rule.Applies(6) != true {
		t.Fatal("Expected true but got false instead")
	}
}

func TestSingleModuloRuleAppliesReturnsFalseIfNotDivisible(t *testing.T) {

}
