package main_test

import (
	"MathExercisesGenerator/docs/fractions"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	f1 := fractions.Create(4, 5)
	f2 := fractions.Create(2, 15)

	summ := fractions.Summ(f1, f2)

	fmt.Println(summ.String())
}
