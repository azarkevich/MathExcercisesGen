package main_test

import (
	"MathExercisesGenerator/docs/fractions"
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

func TestNumeratorZero(t *testing.T) {

	f := fractions.Create(0, 2)

	assert.Equal(t, "0", f.String())
}

func TestNormalization(t *testing.T) {

	tests := []struct {
		name        string
		numerator   int
		denominator int
		expected    string
	}{
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

func TestSumm(t *testing.T) {

	// TODO:
}
