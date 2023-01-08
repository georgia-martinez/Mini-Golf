package main

import (
	// "fmt"
	"math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	x float64
	y float64
	radius float64
	img *ebiten.Image

	vx0 float64
	vy0 float64

	isMoving bool
} 

func (ball *Ball) SetPosition(x float64, y float64) {
	ball.x = x
	ball.y = y
}

func (ball *Ball) SetInitialVelocity(angle float64, power float64) {
	ball.vx0 = math.Cos(angle) * power
	ball.vy0 = math.Sin(angle) * power
}

func (ball *Ball) ResetVelocity() {
	ball.vx0 = 0
	ball.vy0 = 0
}

func (ball *Ball) MovePosition() {
	ball.vx0 *= .8
	ball.vy0 *= .8

	ball.x += ball.vx0
	ball.y += ball.vy0

	if(math.Round(ball.vx0) == 0 && math.Round(ball.vy0) == 0) {
		ball.ResetVelocity()
		ball.isMoving = false
	}

	// fmt.Println(ball)
}