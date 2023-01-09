package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	x      float64
	y      float64
	radius float64
	img    *ebiten.Image

	centerX float64
	centerY float64

	vx0 float64
	vy0 float64

	canStartMoving bool
	isMoving       bool
}

func (ball *Ball) SetPosition(x float64, y float64) {
	ball.x = x
	ball.y = y

	var width, height = ball.img.Size()

	ball.centerX = ball.x + (float64(width) / 2)
	ball.centerY = ball.y + (float64(height) / 2)
}

func (ball *Ball) SetInitialVelocity(angle float64, power float64) {
	ball.vx0 = -math.Cos(angle) * power
	ball.vy0 = math.Sin(angle) * power
}

func (ball *Ball) ResetVelocity() {
	ball.vx0 = 0
	ball.vy0 = 0

	ball.isMoving = false
}

func (ball *Ball) MovePosition() {
	ball.SetPosition(ball.x+ball.vx0, ball.y+ball.vy0)

	ball.vx0 *= .9
	ball.vy0 *= .9

	if math.Round(ball.vx0) == 0 && math.Round(ball.vy0) == 0 {
		ball.ResetVelocity()
	}
}

func (ball *Ball) IsClicked(mouseX float64, mouseY float64) bool {
	return GetDistance(ball.centerX, ball.centerY, mouseX, mouseY) <= ball.radius
}
