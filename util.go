package main

import (
 	"math"
)

func GetDistance(x0 float64, y0 float64, x1 float64, y1 float64) float64 {
	return math.Sqrt(math.Pow(x1 - x0, 2) + math.Pow(y1 - y0, 2))
}