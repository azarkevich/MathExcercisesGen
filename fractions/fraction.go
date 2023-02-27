package fractions

import (
	"fmt"
	"strings"
)

type Fraction struct {
	isNegative  bool
	integerPart int
	numerator   int
	denominator int
}

var Zero Fraction = Fraction{
	isNegative:  false,
	integerPart: 0,
	numerator:   0,
	denominator: 1,
}

func (f Fraction) IsNegative() bool {
	return f.isNegative
}

func (f Fraction) IntegerPart() int {
	return f.integerPart
}

func (f Fraction) HasFractionPart() bool {
	return f.numerator != 0
}

func (f Fraction) Numerator() int {
	return f.numerator
}

func (f Fraction) Denominator() int {
	return f.denominator
}

func Create(numerator int, denominator int) Fraction {
	if denominator == 0 {
		panic("divide by zero.")
	}
	return Fraction{numerator: numerator, denominator: denominator}.normalize()
}

type DenormalizedFraction struct {
	Numerator   int
	Denominator int
}

func gcd(num1, num2 int) int {
	gcd := num1
	if gcd > num2 {
		gcd = num2
	}

	for gcd > 1 {

		if num1%gcd == 0 && num2%gcd == 0 {
			return gcd
		}
		gcd--
	}

	// should be 1
	return gcd
}

func (f Fraction) normalize() (ret Fraction) {

	ret.isNegative = f.isNegative

	ret.numerator = f.numerator
	if ret.numerator < 0 {
		ret.numerator = -ret.numerator
		ret.isNegative = !ret.isNegative
	}

	ret.integerPart = f.integerPart
	if ret.integerPart < 0 {
		ret.integerPart = -ret.integerPart
		ret.isNegative = !ret.isNegative
	}

	ret.denominator = f.denominator
	if ret.denominator < 0 {
		ret.denominator = -ret.denominator
		ret.isNegative = !ret.isNegative
	}

	for ret.numerator >= ret.denominator {

		ret.numerator -= ret.denominator
		ret.integerPart++
	}

	// check if can be simplified
	gcd := gcd(ret.numerator, ret.denominator)
	if gcd > 1 {
		ret.numerator /= gcd
		ret.denominator /= gcd
	}

	if ret.numerator == 0 && ret.integerPart == 0 {
		ret = Zero
	}

	return
}

func (f Fraction) denormalize() (numerator int, denominator int) {

	denominator = f.denominator
	numerator = f.numerator
	for f.integerPart > 0 {
		numerator += denominator
		f.integerPart--
	}

	if f.isNegative {
		numerator = -numerator
	}

	return
}

func lcm(num1, num2 int) (lcm int) {

	// find in bruteforce
	lcm = num1
	if lcm < num2 {
		lcm = num2
	}

	for {

		if lcm%num1 == 0 && lcm%num2 == 0 {
			return
		}

		lcm++
	}
}

func (f Fraction) Abs() Fraction {

	if f.IsNegative() {
		return Fraction{
			isNegative:  false,
			integerPart: f.integerPart,
			numerator:   f.numerator,
			denominator: f.denominator,
		}
	}

	return f
}

func (f Fraction) Neg() Fraction {

	if f == Zero {
		return f
	}

	return Fraction{
		isNegative:  !f.isNegative,
		integerPart: f.integerPart,
		numerator:   f.numerator,
		denominator: f.denominator,
	}

}

func (f Fraction) Add(other Fraction) Fraction {

	numerator, denominator := f.denormalize()
	otherNumerator, otherDenominator := other.denormalize()

	// bring to common denominator
	lcm := lcm(denominator, otherDenominator)

	numerator = numerator * (lcm / denominator)
	otherNumerator = otherNumerator * (lcm / otherDenominator)

	ret := Create(numerator+otherNumerator, lcm)

	return ret
}

func (f Fraction) Sub(other Fraction) Fraction {

	return f.Add(other.Neg())

}

func (f Fraction) Reciprocal() Fraction {

	numerator, denominator := f.denormalize()
	return Create(denominator, numerator)

}

func (f Fraction) Mul(other Fraction) Fraction {

	numerator, denominator := f.denormalize()
	otherNumerator, otherDenominator := other.denormalize()

	return Create(numerator*otherNumerator, denominator*otherDenominator)
}

func (f Fraction) Div(other Fraction) Fraction {

	return f.Mul(other.Reciprocal())

}

// Stringer interface
func (f Fraction) String() string {
	b := strings.Builder{}
	if f.IsNegative() {
		b.WriteString("-")
	}
	hasFractionPart := f.HasFractionPart()
	hasIntegerPart := f.IntegerPart() > 0

	if hasIntegerPart || !hasFractionPart {
		b.WriteString(fmt.Sprintf("%d", f.IntegerPart()))
	}

	if hasIntegerPart && hasFractionPart {
		b.WriteString("*")
	}

	if hasFractionPart {
		if hasIntegerPart {
			b.WriteString("(")
		}
		b.WriteString(fmt.Sprintf("%d/%d", f.Numerator(), f.Denominator()))
		if hasIntegerPart {
			b.WriteString(")")
		}
	}
	return b.String()
}
