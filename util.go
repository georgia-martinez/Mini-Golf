package main

import (
 	"math"
)

func GetDistance(x0 float64, y0 float64, x1 float64, y1 float64) float64 {
	return math.Sqrt(math.Pow(x1 - x0, 2) + math.Pow(y1 - y0, 2))
}

func GetClosestNumber(a float64, b float64, c float64) float64 {
	if math.Abs(c - a) < math.Abs(c - b) {
		return a
	} else {
        return b
	}
}