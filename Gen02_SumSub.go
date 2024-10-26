package main

import (
	"fmt"
	"math/rand"
	"time"
)

// LaTeX to PDF with https://www.overleaf.com/

type Gen02_SumSubOptionsStruct struct {
	WithAnswers    bool
	RandomSeed     int
	EquationsCount int
	MaxResult      int
	Columns        int
}

var Gen02_SumSubOptions = Gen02_SumSubOptionsStruct{
	WithAnswers:    false,
	RandomSeed:     3,
	EquationsCount: 50,
	MaxResult:      30,
	Columns:        3,
}

type equation struct {
	First  int
	Sign   string
	Second int
	Result int
}

func generate(equations []equation, r *rand.Rand) []equation {

	for len(equations) < cap(equations) {

		first := r.Intn(Gen02_SumSubOptions.MaxResult)
		second := r.Intn(Gen02_SumSubOptions.MaxResult)

		// too easy equation
		if first == 0 || second == 0 {
			continue
		}

		var op string
		var result int
		// random operation
		switch r.Intn(2) {
		case 0:
			op = "-"
			result = first - second
		case 1:
			op = "+"
			result = first + second
		}

		// too complex equation
		if result > Gen02_SumSubOptions.MaxResult || result < 0 {
			continue
		}

		newEq := equation{
			First:  first,
			Sign:   op,
			Second: second,
			Result: result,
		}

		equations = append(equations, newEq)
	}
	return equations
}

func Gen02_SumSub() {

	seed := int64(Gen02_SumSubOptions.RandomSeed)
	if seed == 0 {
		// div by 1000 for reduce seed number size
		seed = int64(time.Now().UTC().Unix() / 1000)
	}

	r := rand.New(rand.NewSource(seed))

	equations := make([]equation, 0, Gen02_SumSubOptions.EquationsCount)

	equations = generate(equations, r)

	fmt.Println("\\documentclass{article}")
	fmt.Println("\\usepackage[margin=0.5in]{geometry}")
	fmt.Println("\\usepackage{setspace}")
	fmt.Println("\\usepackage{multicol}")
	fmt.Println("\\begin{document}")
	fmt.Println("\\LARGE")

	fmt.Printf("{\\small Seed: %d Count:%d}\n", seed, Gen02_SumSubOptions.EquationsCount)
	fmt.Println("\\newline")
	fmt.Println("")

	fmt.Printf("\\begin{multicols}{%v}\n", Gen02_SumSubOptions.Columns)

	for i, s := range equations {

		// draw equation
		fmt.Printf("{\\scriptsize\\textbf %d)} ", i+1)
		fmt.Print("$")

		fmt.Printf("%v %v %v = ", s.First, s.Sign, s.Second)

		if Gen02_SumSubOptions.WithAnswers {
			fmt.Printf("%v", s.Result)
		}

		fmt.Println("$")

		fmt.Println("\\vskip 0.1in")
	}

	fmt.Println("\\end{multicols}")
	fmt.Println("\\end{document}")
}
