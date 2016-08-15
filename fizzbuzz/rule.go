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
