package main

import (
	"math"
)

func Calc(x float64) float64 {
		x += math.Sqrt(x)
		return x
}
