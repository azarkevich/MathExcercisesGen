package main

import (
	"MathExercisesGenerator/fractions"
	"fmt"
	"math/rand"
	"time"
)

// LaTeX to PDF with https://www.overleaf.com/

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
	ShowSolutionInPlace:    true,
	EquationsCount:         6,
	RandomSeed:             1,
}

func drawFraction(f fractions.Fraction) {
	if f.IsNegative() {
		fmt.Print(" - ")
	} else {
		fmt.Print(" + ")
	}

	drawFractionNoSign(f)
}

func drawSoleFraction(f fractions.Fraction) {
	if f.IsNegative() {
		fmt.Print(" - ")
	}

	drawFractionNoSign(f)
}

func drawFractionNoSign(f fractions.Fraction) {

	hasFractionPart := f.HasFractionPart()
	if f.IntegerPart() > 0 || !hasFractionPart {
		fmt.Printf("%d", f.IntegerPart())
	}

	if hasFractionPart {
		fmt.Printf("\\frac{%d}{%d}", f.Numerator(), f.Denominator())
	}
}

func Gen01_Fractions() {

	_ = fractions.Fraction{}

	seed := int64(Gen01_FractionsOptions.RandomSeed)
	if seed == 0 {
		// div by 1000 for reduce seed number size
		seed = int64(time.Now().UTC().Unix() / 1000)
	}

	rand.Seed(seed)

	solutions := make([]fractions.Fraction, Gen01_FractionsOptions.EquationsCount)

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

		equation := make([]fractions.Fraction, fractionsCount)

		for f := 0; f < len(equation); f++ {
			numerator := Gen01_FractionsOptions.MinNumerator + rand.Intn(Gen01_FractionsOptions.MaxNumerator-Gen01_FractionsOptions.MinNumerator)
			denominator := Gen01_FractionsOptions.MinDenominator + rand.Intn(Gen01_FractionsOptions.MaxDenominator-Gen01_FractionsOptions.MinDenominator)

			if f > 0 && rand.Intn(2) == 1 {
				numerator = -numerator
			}

			fraction := fractions.Create(numerator, denominator)

			equation[f] = fraction
		}

		// evaluate

		solution := equation[0]
		for _, other := range equation[1:] {
			solution = solution.Add(other)
		}

		// too complex solution
		if solution.Denominator() > Gen01_FractionsOptions.MaxSolutionDenominator {
			continue
		}

		if solution.IsNegative() && Gen01_FractionsOptions.NoNegativeResults {
			// skip this equation, because it generate negative result
			continue
		}

		// draw equation
		fmt.Printf("{\\small\\textbf %d.} ", eqInd+1)
		fmt.Print("$")

		for f := 0; f < len(equation); f++ {
			if f == 0 {
				drawSoleFraction(equation[f])
			} else {
				drawFraction(equation[f])
			}
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

	if !Gen01_FractionsOptions.ShowSolutionInPlace {

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
