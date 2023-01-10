package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

// import (
// 	"math"
// 	"github.com/hajimehoshi/ebiten/v2"
// )

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

func (wall *Wall) GetTag() string {
	return "wall"
}

func (wall *Wall) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, wall.x, wall.y, wall.height, wall.width, color.Gray{ 169 })
}