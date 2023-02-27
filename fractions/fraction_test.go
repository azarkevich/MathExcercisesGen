package fractions_test

import (
	"MathExercisesGenerator/fractions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDenominatorZero(t *testing.T) {

	defer func() {
		recover()
	}()

	fractions.Create(0, 0)

	assert.Fail(t, "Should panic on denominator == 0")
}

func TestNormalization(t *testing.T) {

	tests := []struct {
		name        string
		numerator   int
		denominator int
		expected    string
	}{
		{name: "Zero/Number", numerator: 0, denominator: 2, expected: "0"},
		{name: "Zero/-Number", numerator: 0, denominator: -2, expected: "0"},
		{name: "Negative numerator", numerator: -1, denominator: 2, expected: "-1/2"},
		{name: "Negative numerator with integer part", numerator: -3, denominator: 2, expected: "-1*(1/2)"},
		{name: "Simple", numerator: 2, denominator: 3, expected: "2/3"},
		{name: "With integer part", numerator: 7, denominator: 2, expected: "3*(1/2)"},
		{name: "Simplifieble", numerator: 3, denominator: 9, expected: "1/3"},
		{name: "Simplifieble with integer part", numerator: 12, denominator: 9, expected: "1*(1/3)"},
		{name: "Negative simplifieble with integer part", numerator: -12, denominator: 9, expected: "-1*(1/3)"},
		{name: "Just integer", numerator: 1, denominator: 1, expected: "1"},
		{name: "Just integer, negative", numerator: -1, denominator: 1, expected: "-1"},
		{name: "Just integer, simplifiable", numerator: 3, denominator: 3, expected: "1"},
		{name: "Just integer, simplifiable, negative", numerator: -3, denominator: 3, expected: "-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fractions.Create(tt.numerator, tt.denominator)
			assert.Equal(t, tt.expected, f.String())
		})
	}
}

func TestAbs(t *testing.T) {

	tests := []struct {
		name     string
		fract    fractions.Fraction
		expected string
	}{
		{name: "Abs(zero)", fract: fractions.Zero, expected: "0"},
		{name: "Abs(num)", fract: fractions.Create(1, 2), expected: "1/2"},
		{name: "Abs(num2)", fract: fractions.Create(3, 2), expected: "1*(1/2)"},
		{name: "Abs(-num)", fract: fractions.Create(-1, 2), expected: "1/2"},
		{name: "Abs(-num2)", fract: fractions.Create(-3, 2), expected: "1*(1/2)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.fract.Abs()
			assert.Equal(t, tt.expected, f.String())
		})
	}
}

func TestNeg(t *testing.T) {

	tests := []struct {
		name     string
		fract    fractions.Fraction
		expected string
	}{
		{name: "Neg(zero)", fract: fractions.Zero, expected: "0"},
		{name: "Neg(num)", fract: fractions.Create(1, 2), expected: "-1/2"},
		{name: "Neg(num2)", fract: fractions.Create(3, 2), expected: "-1*(1/2)"},
		{name: "Neg(-num)", fract: fractions.Create(-1, 2), expected: "1/2"},
		{name: "Neg(-num2)", fract: fractions.Create(-3, 2), expected: "1*(1/2)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.fract.Neg()
			assert.Equal(t, tt.expected, f.String())
		})
	}
}

func TestAdd(t *testing.T) {

	tests := []struct {
		name     string
		first    fractions.Fraction
		second   fractions.Fraction
		expected string
	}{
		{name: "Add zeros", first: fractions.Zero, second: fractions.Create(0, 1), expected: "0"},
		{name: "Add zero+num", first: fractions.Zero, second: fractions.Create(1, 2), expected: "1/2"},
		{name: "Add zero+num2", first: fractions.Zero, second: fractions.Create(3, 2), expected: "1*(1/2)"},
		{name: "Add num+zero", first: fractions.Create(1, 2), second: fractions.Zero, expected: "1/2"},
		{name: "Add num2+zero", first: fractions.Create(3, 2), second: fractions.Zero, expected: "1*(1/2)"},
		{name: "Add num + num", first: fractions.Create(1, 2), second: fractions.Create(2, 3), expected: "1*(1/6)"},
		{name: "Add num + -num", first: fractions.Create(1, 2), second: fractions.Create(-2, 3), expected: "-1/6"},
		{name: "Add -num + num", first: fractions.Create(-1, 2), second: fractions.Create(2, 3), expected: "1/6"},
		{name: "Add -num + -num", first: fractions.Create(-1, 2), second: fractions.Create(-2, 3), expected: "-1*(1/6)"},
		{name: "Add X + -X", first: fractions.Create(1, 2), second: fractions.Create(-1, 2), expected: "0"},
		{name: "Add -X + X", first: fractions.Create(-1, 2), second: fractions.Create(1, 2), expected: "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.first.Add(tt.second)
			assert.Equal(t, tt.expected, f.String())
		})
	}
}

func TestSub(t *testing.T) {

	tests := []struct {
		name     string
		first    fractions.Fraction
		second   fractions.Fraction
		expected string
	}{
		{name: "Sub (num - num)", first: fractions.Create(1, 2), second: fractions.Create(2, 3), expected: "-1/6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.first.Sub(tt.second)
			assert.Equal(t, tt.expected, f.String())
		})
	}
}

func TestReciprocalZero(t *testing.T) {

	defer func() {
		recover()
	}()

	fractions.Zero.Reciprocal()

	assert.Fail(t, "Should panic on reciprocation 0/1")
}

func TestReciprocal(t *testing.T) {

	tests := []struct {
		name     string
		fract    fractions.Fraction
		expected string
	}{
		{name: "Reciprocal(num)", fract: fractions.Create(1, 2), expected: "2"},
		{name: "Reciprocal(num2)", fract: fractions.Create(2, 3), expected: "1*(1/2)"},
		{name: "Reciprocal(num3)", fract: fractions.Create(3, 2), expected: "2/3"},
		{name: "Reciprocal(-num)", fract: fractions.Create(-1, 2), expected: "-2"},
		{name: "Reciprocal(-num2)", fract: fractions.Create(-2, 3), expected: "-1*(1/2)"},
		{name: "Reciprocal(-num3)", fract: fractions.Create(-3, 2), expected: "-2/3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.fract.Reciprocal()
			assert.Equal(t, tt.expected, f.String())
		})
	}
}

func TestMul(t *testing.T) {

	tests := []struct {
		name     string
		first    fractions.Fraction
		second   fractions.Fraction
		expected string
	}{
		{name: "Mul (num * num)", first: fractions.Create(1, 2), second: fractions.Create(2, 3), expected: "1/3"},
		{name: "Mul (0 * num)", first: fractions.Zero, second: fractions.Create(2, 3), expected: "0"},
		{name: "Mul (num * 0)", first: fractions.Create(1, 2), second: fractions.Zero, expected: "0"},
		{name: "Mul (num * 1)", first: fractions.Create(1, 2), second: fractions.Create(1, 1), expected: "1/2"},
		{name: "Mul (1 * num)", first: fractions.Create(1, 1), second: fractions.Create(2, 3), expected: "2/3"},
		{name: "Mul (2 * 3)", first: fractions.Create(2, 1), second: fractions.Create(3, 1), expected: "6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.first.Mul(tt.second)
			assert.Equal(t, tt.expected, f.String())
		})
	}
}

func TestDivZero(t *testing.T) {

	defer func() {
		recover()
	}()

	fractions.Create(1, 2).Div(fractions.Zero)

	assert.Fail(t, "Should panic on divide by zero fraction")
}

func TestDiv(t *testing.T) {

	tests := []struct {
		name     string
		first    fractions.Fraction
		second   fractions.Fraction
		expected string
	}{
		{name: "Div (num / num)", first: fractions.Create(1, 2), second: fractions.Create(2, 3), expected: "3/4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.first.Div(tt.second)
			assert.Equal(t, tt.expected, f.String())
		})
	}
}
