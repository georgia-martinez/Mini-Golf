package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject interface {
	Draw(screen *ebiten.Image)
	GetTag() string
}