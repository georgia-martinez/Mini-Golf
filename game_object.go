package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject interface {
	Draw(screen *ebiten.Image)
	Tag() string
	Top() float64
	Bottom() float64
	Left() float64
	Right() float64
}
