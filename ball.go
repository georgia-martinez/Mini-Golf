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

	dirX float64
	dirY float64
}

func (ball *Ball) SetPosition(x float64, y float64) {
	ball.x = x
	ball.y = y

	width, height := ball.img.Size()

	ball.centerX = ball.x + (float64(width) / 2)
	ball.centerY = ball.y + (float64(height) / 2)
}

func (ball *Ball) SetInitialVelocity(angle float64, power float64) {
	ball.vx0 = -math.Cos(angle) * power
	ball.vy0 = math.Sin(angle) * power

	ball.dirX = 1
	ball.dirY = 1
}

func (ball *Ball) ResetVelocity() {
	ball.vx0 = 0
	ball.vy0 = 0

	ball.isMoving = false
}

func (ball *Ball) MovePosition() {

	// Changing the ball's velocity
	const DRAG = .9

	ball.vx0 *= ball.dirX * DRAG
	ball.vy0 *= ball.dirY * DRAG

	// Make ball bounce if it hits a wall
	width, height := ball.img.Size()
	
	w := float64(width) / 2
	h := float64(height) / 2

	if ball.centerX - w < 0 || ball.centerX + w > screenWidth {
		ball.dirX *= -1
		ball.x = GetClosestNumber(0, float64(screenWidth - width), ball.centerX)
	}

	if ball.centerY - h < 0 || ball.centerY + h > screenHeight {
		ball.dirY *= -1
		ball.y = GetClosestNumber(0, float64(screenHeight - height), ball.centerY)
	}

	// Changing the ball's position
	ball.SetPosition(ball.x+ball.vx0, ball.y+ball.vy0)

	// Reset the velocity is the ball is done moving
	if math.Round(ball.vx0) == 0 && math.Round(ball.vy0) == 0 {
		ball.ResetVelocity()
	}
}

func (ball *Ball) IsClicked(mouseX float64, mouseY float64) bool {
	return GetDistance(ball.centerX, ball.centerY, mouseX, mouseY) <= ball.radius
}
