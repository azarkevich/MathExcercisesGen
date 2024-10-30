package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"math/rand"
	"strings"
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
	RandomSeed:     4,
	EquationsCount: 69,	// 69 fit in one page
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
		if result < 0 || result > Gen02_SumSubOptions.MaxResult {
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

	var sb strings.Builder

	sb.WriteString("\\documentclass{article}")
	sb.WriteString("\\usepackage[margin=0.5in]{geometry}")
	sb.WriteString("\\usepackage{setspace}")
	sb.WriteString("\\usepackage{multicol}")
	sb.WriteString("\\begin{document}")
	sb.WriteString("\\LARGE")

	sb.WriteString(fmt.Sprintf("{\\small Seed: %d Count:%d}\n", seed, Gen02_SumSubOptions.EquationsCount))
	sb.WriteString("\\newline")
	sb.WriteString("")

	sb.WriteString(fmt.Sprintf("\\begin{multicols}{%v}\n", Gen02_SumSubOptions.Columns))

	for i, s := range equations {

		// draw equation
		sb.WriteString(fmt.Sprintf("{\\scriptsize\\textbf %d)} ", i+1))
		sb.WriteString("$")

		sb.WriteString(fmt.Sprintf("%v %v %v = ", s.First, s.Sign, s.Second))

		if Gen02_SumSubOptions.WithAnswers {
			sb.WriteString(fmt.Sprintf("%v", s.Result))
		}

		sb.WriteString("$")

		sb.WriteString("\\vskip 0.1in")
	}

	sb.WriteString("\\end{multicols}")
	sb.WriteString("\\end{document}")

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	clipboard.Write(clipboard.FmtText, []byte(sb.String()))

	fmt.Println("Copied to clipboard!")
}
