package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	x, y, radius float64
	img *ebiten.Image
} 

func (ball Ball) SetPosition(x float64, y float64) {
	ball.x = x
	ball.y = y
}