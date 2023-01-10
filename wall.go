package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type Wall struct {
	x float64
	y float64
	width float64
	height float64
}

func NewWall(x float64, y float64, width float64, height float64) *Wall {
	return &Wall {
		x,
		y,
		width,
		height,
	}
}

// GameObject interface methods

func (wall *Wall) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, wall.x, wall.y, wall.height, wall.width, color.Gray{ 169 })
}

func (wall *Wall) Tag() string {
	return "wall"
}

func (wall *Wall) Top() float64 {
	return wall.y
}

func (wall *Wall) Bottom() float64 {
	return wall.y + wall.height
}

func (wall *Wall) Left() float64 {
	return wall.x
}

func (wall *Wall) Right() float64 {
	return wall.x + wall.width
}
