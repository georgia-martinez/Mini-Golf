package main

import (
    "math"
)

type MouseEvent struct {
	x0, y0 float64 // initial position when mouse is first pressed
}

func (m MouseEvent) GetDistance(x1 float64, y1 float64) float64 {
	return math.Sqrt(math.Pow(x1 - m.x0, 2) + math.Pow(y1 - m.y0, 2))
}