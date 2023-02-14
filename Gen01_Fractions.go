package main

import (
	"fmt"
	"math/rand"
	"time"
)

// LaTeX to PDF with https://www.overleaf.com/

type Fraction struct {
	Numerator   int
	Denominator int
}

type Gen01_FractionsOptionsStruct struct {
	NoNegativeResults      bool
	MaxSolutionDenominator int
	MinPartiesCount        int
	MaxPartiesCount        int
	MinNumerator           int
	MaxNumerator           int
	MinDenominator         int
	MaxDenominator         int
	ShowSolutionInPlace    bool
	EquationsCount         int
	RandomSeed             int
}

var Gen01_FractionsOptions = Gen01_FractionsOptionsStruct{
	NoNegativeResults:      true,
	MaxSolutionDenominator: 50,
	MinPartiesCount:        2,
	MaxPartiesCount:        3,
	MinNumerator:           1,
	MaxNumerator:           17,
	MinDenominator:         2,
	MaxDenominator:         23,
	ShowSolutionInPlace:    false,
	EquationsCount:         42,
	RandomSeed:             1,
}

func drawSoleFraction(f Fraction) {
	if f.Numerator < 0 {
		fmt.Print(" - ")
	}

	drawFractionNoSign(f)
}

func drawFractionNoSign(f Fraction) {

	normNumerator := f.Numerator
	if normNumerator < 0 {
		normNumerator = -normNumerator
	}

	if normNumerator > f.Denominator {
		intPart := normNumerator / f.Denominator

		fmt.Printf("%d", intPart)

		normNumerator = normNumerator - intPart*f.Denominator
	}

	if normNumerator > 0 {
		fmt.Printf("\\frac{%d}{%d}", normNumerator, f.Denominator)
	}
}

func lcm(nums []int) int {
	// find in bruteforce
	lcm := nums[0]
	for {
		notLcm := false
		for i := 0; i < len(nums); i++ {
			if lcm < nums[i] || lcm%nums[i] != 0 {
				notLcm = true
				break
			}
		}

		if !notLcm {
			return lcm
		}
		lcm++
	}
}

func NormalizeFraction(f Fraction) Fraction {
	max := f.Numerator * f.Denominator
	if max < 0 {
		max = -max
	}
	min := f.Numerator
	if min < 0 {
		min = -min
	}
	if min < f.Denominator {
		min = f.Denominator
	}
	for i := max; i >= min; i-- {
		if f.Numerator%i == 0 && f.Denominator%i == 0 {
			return Fraction{
				Numerator:   f.Numerator / i,
				Denominator: f.Denominator / i,
			}
		}
	}

	return f
}

func solveEquation(fractions []Fraction) Fraction {
	// find min
	denominators := make([]int, len(fractions))
	for i := 0; i < len(fractions); i++ {
		denominators[i] = fractions[i].Denominator
	}

	lcm := lcm(denominators)

	//fmt.Println("% lcm = ", lcm)

	solution := Fraction{
		Numerator:   0,
		Denominator: lcm,
	}

	for i := 0; i < len(fractions); i++ {
		mul := lcm / fractions[i].Denominator
		solution.Numerator += fractions[i].Numerator * mul
	}

	return NormalizeFraction(solution)
}

func Gen01_Fractions() {

	seed := int64(Gen01_FractionsOptions.RandomSeed)
	if seed == 0 {
		// div by 1000 for reduce seed number size
		seed = int64(time.Now().UTC().Unix() / 1000)
	}

	rand.Seed(seed)

	solutions := make([]Fraction, Gen01_FractionsOptions.EquationsCount)

	fmt.Println("\\documentclass{article}")
	fmt.Println("\\usepackage[margin=0.5in]{geometry}")
	fmt.Println("\\usepackage{setspace}")
	fmt.Println("\\begin{document}")
	fmt.Println("\\huge")

	fmt.Printf("{\\small Seed: %d Count:%d}\n", seed, Gen01_FractionsOptions.EquationsCount)
	fmt.Println("\\newline\\newline")
	fmt.Println("")

	eqInd := 0

	for eqInd < Gen01_FractionsOptions.EquationsCount {

		fractionsCount := Gen01_FractionsOptions.MinPartiesCount

		if Gen01_FractionsOptions.MaxPartiesCount > Gen01_FractionsOptions.MinPartiesCount {
			fractionsCount = Gen01_FractionsOptions.MinPartiesCount + rand.Intn(Gen01_FractionsOptions.MaxPartiesCount-Gen01_FractionsOptions.MinPartiesCount+1)
		}

		fractions := make([]Fraction, fractionsCount)

		for f := 0; f < len(fractions); f++ {
			fractions[f] = Fraction{
				Numerator:   Gen01_FractionsOptions.MinNumerator + rand.Intn(Gen01_FractionsOptions.MaxNumerator-Gen01_FractionsOptions.MinNumerator),
				Denominator: Gen01_FractionsOptions.MinDenominator + rand.Intn(Gen01_FractionsOptions.MaxDenominator-Gen01_FractionsOptions.MinDenominator),
			}

			if f > 0 && rand.Intn(2) == 1 {
				fractions[f].Numerator = -fractions[f].Numerator
			}

			fractions[f] = NormalizeFraction(fractions[f])

			if fractions[f].Denominator == 1 {
				fractions = nil
				break
			}
		}

		if fractions == nil {
			continue
		}

		// validate
		solution := solveEquation(fractions)

		if solution.Denominator > Gen01_FractionsOptions.MaxSolutionDenominator {
			continue
		}

		if solution.Numerator < 0 && Gen01_FractionsOptions.NoNegativeResults {
			// skip this equation, because it generate negative result
			continue
		}

		// draw equation
		fmt.Printf("{\\small\\textbf %d.} ", eqInd+1)
		fmt.Print("$")

		for f := 0; f < len(fractions); f++ {
			if fractions[f].Numerator < 0 {
				fmt.Print(" - ")
			} else if f > 0 {
				fmt.Print(" + ")
			}

			drawFractionNoSign(fractions[f])
		}
		fmt.Print(" = ")

		solutions[eqInd] = solution

		if Gen01_FractionsOptions.ShowSolutionInPlace {
			drawSoleFraction(solution)
		}
		fmt.Println("$")

		fmt.Println("\\newline")
		fmt.Println("")

		eqInd++
	}

	if Gen01_FractionsOptions.ShowSolutionInPlace == false {

		fmt.Println("\\newline\\newline\\newline\\newline")
		fmt.Println("")
		fmt.Println("Solutions:")
		fmt.Println("\\doublespacing")
		fmt.Println("")

		for i := 0; i < len(solutions); i++ {
			fmt.Printf("{\\small\\textbf %d.} $", i+1)
			drawSoleFraction(solutions[i])
			fmt.Println("$\\quad")
		}
	}

	fmt.Println("\\end{document}")

}
