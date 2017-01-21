// `tempconv` is supposed to be a standalone package, so this `main` function
// shouldn't be part of part of it.
package main

import (
	"fmt"

	"github.com/lmbarros/go_playground/gopl_exercises/ex_2_1-tempconv/tempconv"
)

func main() {
	cTemps := []tempconv.Celsius{tempconv.AbsoluteZeroC, -12.0, tempconv.FreezingC, 11.1, 35.3, tempconv.BoilingC}
	fTemps := []tempconv.Fahrenheit{-30.2, 0.0, 25.5, 48.7, 102.9, 222.2}
	kTemps := []tempconv.Kelvin{0.0, 50.1, 333.2, 489.1, 1000.0, 1234.5}

	for _, c := range cTemps {
		fmt.Printf("%v in is %v\n", c, tempconv.CToK(c))
	}

	for _, f := range fTemps {
		fmt.Printf("%v in is %v\n", f, tempconv.FToK(f))
	}

	for _, k := range kTemps {
		fmt.Printf("%v in is %v (AKA %v)\n", k, tempconv.KToC(k), tempconv.KToF(k))
	}
}
