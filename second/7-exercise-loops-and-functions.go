package main

import (
	"fmt"
	"math"
)

var diff float64 = 10.0

const diffLim float64 = 0.001

func Sqrt(x float64) (float64, int) {
	z := 1.0
	i := 0
	for i < 10 && diff > diffLim {
		u := z - (z*z-x)/(2*z)
		diff = math.Abs(z - u)
		z = u
		i++
	}
	return z, i
}

func main() {
	z, i := Sqrt(2)
	fmt.Printf("The approximation %v with precision less than %v (%v) was obtained in %v iterations.", z, diffLim, diff, i)
}
