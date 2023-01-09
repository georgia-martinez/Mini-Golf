package main

import (
	// "fmt"
    "math"
)

type MouseEvent struct {
	x0, y0 float64 // initial position when mouse is first pressed
	power float64
	angle float64
}

func (m *MouseEvent) SetPower(x1 float64, y1 float64) {
	m.power = math.Sqrt(math.Pow(x1 - m.x0, 2) + math.Pow(y1 - m.y0, 2)) / 4.5
}

func (m *MouseEvent) SetAngle(x1 float64, y1 float64) {
	m.angle = -math.Atan2(y1 - m.y0, x1 - m.x0)
}

