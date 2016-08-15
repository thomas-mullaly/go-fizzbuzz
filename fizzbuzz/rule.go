package fizzbuzz

import "strconv"

// Rule defines what a fizz buzz rule should look like
type Rule interface {
	Applies(i int) bool
	Result(i int) string
}

// SingleModuloRule represents a rule which applies if a number is divisble by a single value.
type SingleModuloRule struct {
	moduloValue int
	result      string
}

// NewSingleModuloRule returns a new rule using the moduloValue and result string
func NewSingleModuloRule(moduloValue int, result string) *SingleModuloRule {
	return &SingleModuloRule{
		moduloValue: moduloValue,
		result:      result,
	}
}

// Applies returns true if i is divisble by the rules moduloValue
func (r *SingleModuloRule) Applies(i int) bool {
	return i%r.moduloValue == 0
}

// Result returns the result of the rule
func (r *SingleModuloRule) Result(i int) string {
	return r.result
}

// DoubleModuloRule represents a FizzBuzz rule that only applies when the value is divisible by
// two denominators
type DoubleModuloRule struct {
	moduloValue1 int
	moduloValue2 int
	result       string
}

// NewDoubleModuloRule creates a new DoubleModuloRule using the given parameters
func NewDoubleModuloRule(moduloValue1 int, moduloValue2 int, result string) *DoubleModuloRule {
	return &DoubleModuloRule{
		moduloValue1: moduloValue1,
		moduloValue2: moduloValue2,
		result:       result,
	}
}

// Applies returns true if i is divisible by both modulo values
func (r *DoubleModuloRule) Applies(i int) bool {
	return i%r.moduloValue1 == 0 && i%r.moduloValue2 == 0
}

// Result returns the result of the rule
func (r *DoubleModuloRule) Result(i int) string {
	return r.result
}

// NumberEchoRule Applies to all values and returns the value in string form
type NumberEchoRule struct{}

// Applies always returns true
func (r *NumberEchoRule) Applies(i int) bool {
	return true
}

// Result always returns the string representation of i
func (r *NumberEchoRule) Result(i int) string {
	return strconv.Itoa(i)
}
