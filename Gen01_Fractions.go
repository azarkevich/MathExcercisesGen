package main

import (
	"MathExercisesGenerator/fractions"
	"fmt"
	"math/rand"
	"time"
)

// LaTeX to PDF with https://www.overleaf.com/

type Gen01_FractionsOptionsStruct struct {
	MaxSolutionDenominator int
	MinNumerator           int
	MaxNumerator           int
	MinDenominator         int
	MaxDenominator         int
	ShowSolutionInPlace    bool
	EquationsCount         int
	RandomSeed             int
}

var Gen01_FractionsOptions = Gen01_FractionsOptionsStruct{
	MaxSolutionDenominator: 50,
	MinNumerator:           1,
	MaxNumerator:           17,
	MinDenominator:         2,
	MaxDenominator:         23,
	ShowSolutionInPlace:    false,
	EquationsCount:         5,
	RandomSeed:             1,
}

func drawFraction(f fractions.Fraction, negInParens bool) {
	if f.IsNegative() {
		if negInParens {
			fmt.Print(" ( ")
		}
		fmt.Print(" - ")
	}

	drawFractionNoSign(f)

	if f.IsNegative() && negInParens {
		fmt.Print(" ) ")
	}
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

		var createRandomFraction = func() fractions.Fraction {

			numerator := Gen01_FractionsOptions.MinNumerator + rand.Intn(Gen01_FractionsOptions.MaxNumerator-Gen01_FractionsOptions.MinNumerator)
			denominator := Gen01_FractionsOptions.MinDenominator + rand.Intn(Gen01_FractionsOptions.MaxDenominator-Gen01_FractionsOptions.MinDenominator)

			fraction := fractions.Create(numerator, denominator)

			if rand.Intn(2) == 1 {
				fraction = fraction.Neg()
			}

			return fraction
		}

		fraction1 := createRandomFraction()
		fraction2 := createRandomFraction()

		var solution fractions.Fraction

		var op string
		// random operation
		switch rand.Intn(4) {
		case 0:
			op = "+"
			solution = fraction1.Add(fraction2)
		case 1:
			op = "-"
			solution = fraction1.Sub(fraction2)
		case 2:
			op = "*"
			solution = fraction1.Mul(fraction2)
		case 3:
			op = ":"
			// do not allow divide by 0
			for fraction2 == fractions.Zero {
				fraction2 = createRandomFraction()
			}
			solution = fraction1.Div(fraction2)
		}

		// too complex solution
		if solution.Denominator() > Gen01_FractionsOptions.MaxSolutionDenominator {
			continue
		}

		// draw equation
		fmt.Printf("{\\small\\textbf %d)} ", eqInd+1)
		fmt.Print("$")

		drawFraction(fraction1, false)

		if op == "*" {
			fmt.Print(" \\times ")
		} else if op == ":" {
			fmt.Print(" \\div ")
		} else {
			fmt.Printf(" %v ", op)
		}

		drawFraction(fraction2, true)

		fmt.Print(" = ")

		solutions[eqInd] = solution

		if Gen01_FractionsOptions.ShowSolutionInPlace {
			drawFraction(solution, false)
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
			fmt.Printf("{\\small\\textbf %d)} $", i+1)
			drawFraction(solutions[i], false)
			fmt.Println("$\\quad")
		}
	}

	fmt.Println("\\end{document}")

}
