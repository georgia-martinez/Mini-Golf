package main

import (
	"math"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	img    *ebiten.Image
	x      float64
	y      float64
	radius float64

	top    float64
	bottom float64
	left   float64
	right  float64

	width  float64
	height float64

	centerX float64
	centerY float64

	velX float64
	velY float64

	canStartMoving bool
	isMoving       bool

	gameObjects []GameObject // to check collisions with
}

func NewBall(img *ebiten.Image, x float64, y float64) Ball {
	width, height := img.Size()

	ball := Ball{
		img:    img,
		x:      x,
		y:      y,
		radius: float64(width) / 2,
		width:  float64(width),
		height: float64(height),
	}

	ball.SetPosition(x, y)

	return ball
}

func (ball *Ball) SetPosition(x float64, y float64) {
	ball.x = x
	ball.y = y

	ball.centerX = ball.x + (ball.width / 2)
	ball.centerY = ball.y + (ball.height / 2)

	ball.top = ball.y
	ball.bottom = ball.y + ball.height
	ball.left = ball.x
	ball.right = ball.x + ball.width
}

func (ball *Ball) SetInitialVelocity(angle float64, power float64) {
	ball.velX = -math.Cos(angle) * power
	ball.velY = math.Sin(angle) * power
}

func (ball *Ball) ResetVelocity() {
	ball.velX = 0
	ball.velY = 0

	ball.isMoving = false
}

func (ball *Ball) MovePosition() {

	// Changing the ball's velocity
	const DRAG = .95

	ball.velX *= DRAG
	ball.velY *= DRAG

	// Make ball bounce if it hits a wall
	ball.CheckForScreenCollision()

	for _, object := range ball.gameObjects {
		ball.CheckForCollision(object)
	}

	// Changing the ball's position
	ball.SetPosition(ball.x+ball.velX, ball.y+ball.velY)

	// Check if ball is done moving
	if math.Round(ball.velX) == 0 && math.Round(ball.velY) == 0 {
		ball.ResetVelocity()
	}
}

func (ball *Ball) CheckForScreenCollision() {
	w := ball.width / 2
	h := ball.height / 2

	if ball.centerX-w < 0 || ball.centerX+w > screenWidth {
		ball.velX *= -1
		ball.x = GetClosestNumber(0, float64(screenWidth-ball.width), ball.centerX)
	}

	if ball.centerY-h < 0 || ball.centerY+h > screenHeight {
		ball.velY *= -1
		ball.y = GetClosestNumber(0, float64(screenHeight-ball.height), ball.centerY)
	}
}

func (ball *Ball) CheckForCollision(o GameObject) {

	if !ball.IsTouching(o) {
		return
	}

	var tolerance float64 = 25

	fmt.Println("BEFORE: velX: ", ball.velX, " velY: ", ball.velY)

	const BUFF = 1

	if math.Abs(o.Top() - ball.bottom) < tolerance { // Collision with top side of wall
		ball.velY *= -1
		ball.y = o.Top() - ball.height - BUFF 
	} else if math.Abs(o.Bottom() - ball.top) < tolerance { // Collision with bottom side of wall
		ball.velY *= -1
		ball.y = o.Bottom() + BUFF
	} else if math.Abs(o.Left() - ball.right) < tolerance { // Collision with left side of wall
		ball.velX *= -1
		ball.x = o.Left() - ball.width - BUFF 
	} else if math.Abs(o.Right() - ball.left) < tolerance { // Collision with right side of wall
		ball.velX *= -1
		ball.x = o.Right() + BUFF
		return
	}

	fmt.Println("AFTER: velX: ", ball.velX, " velY: ", ball.velY)
}

func (ball *Ball) IsTouching(o GameObject) bool {

	if ball.right < o.Left() || ball.left > o.Right() {
		return false
	}

	if ball.bottom < o.Top() || ball.top > o.Bottom() {
		return false
	}

	return true
}

func (ball *Ball) IsClicked(mouseX float64, mouseY float64) bool {
	return GetDistance(ball.centerX, ball.centerY, mouseX, mouseY) <= ball.radius
}
