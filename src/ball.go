package main

type Ball struct {
	x, y, radius float64
} 

func (ball Ball) SetPosition(x float64, y float64) {
	ball.x = x
	ball.y = y
}