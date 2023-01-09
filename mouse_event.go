package main

import (
	// "fmt"
    "math"
)

type MouseEvent struct {
	mouseX float64
	mouseY float64
	power float64
	angle float64
}

func (m *MouseEvent) SetPower(x1 float64, y1 float64) {
	m.power = GetDistance(m.mouseX, m.mouseY, x1, y1) / 6
}

func (m *MouseEvent) SetAngle(x1 float64, y1 float64) {
	m.angle = -math.Atan2(y1 - m.mouseY, x1 - m.mouseX)
}

