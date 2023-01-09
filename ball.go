package main

import (
	"math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	img    *ebiten.Image
	x      float64
	y      float64
	radius float64
	centerX float64
	centerY float64
	velX float64
	velY float64
	dirX float64
	dirY float64
	canStartMoving bool
	isMoving       bool
}

func (ball *Ball) SetPosition(x float64, y float64) {
	ball.x = x
	ball.y = y

	width, height := ball.img.Size()

	ball.centerX = ball.x + (float64(width) / 2)
	ball.centerY = ball.y + (float64(height) / 2)
}

func (ball *Ball) SetInitialVelocity(angle float64, power float64) {
	ball.velX = -math.Cos(angle) * power
	ball.velY = math.Sin(angle) * power
}

func (ball *Ball) SetInitialDirection() {
	ball.dirX = 1
	ball.dirY = 1
}

func (ball *Ball) ResetVelocity() {
	ball.velX = 0
	ball.velY = 0

	ball.isMoving = false
}

func (ball *Ball) MovePosition() {

	// Changing the ball's velocity
	const DRAG = .9

	ball.velX *= ball.dirX * DRAG
	ball.velY *= ball.dirY * DRAG

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
	ball.SetPosition(ball.x+ball.velX, ball.y+ball.velY)

	// Reset the velocity is the ball is done moving
	if math.Round(ball.velX) == 0 && math.Round(ball.velY) == 0 {
		ball.ResetVelocity()
	}
}

func (ball *Ball) IsClicked(mouseX float64, mouseY float64) bool {
	return GetDistance(ball.centerX, ball.centerY, mouseX, mouseY) <= ball.radius
}
