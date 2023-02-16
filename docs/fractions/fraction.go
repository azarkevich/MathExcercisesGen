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

func Summ(nums ...Fraction) Fraction {

	accum := nums[0]
	for _, other := range nums[1:] {
		accum = accum.Add(other)
	}
	return accum
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
