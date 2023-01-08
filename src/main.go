package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 500
	screenHeight = 500
)

type Game struct{
	backgroundImage *ebiten.Image
	ball Ball
	mouseDown bool
	mouseEvent MouseEvent
}

func NewGame() *Game {

	var ball = Ball { x: 250, y: 250, radius: 10 }

	var backgroundImage = ebiten.NewImage(screenWidth, screenHeight)
	backgroundImage.Fill(color.Black)

	g := &Game { 
		backgroundImage: backgroundImage, 
		ball: ball,
		mouseDown: false,
		mouseEvent: MouseEvent{},
	}

	return g
}

func (g *Game) Update() error {
	if(inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)) {
		var x0, y0 = ebiten.CursorPosition()
		g.mouseEvent = MouseEvent { x0: float64(x0), y0: float64(y0) }

	} else if (inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)) {
		var x1, y1 = ebiten.CursorPosition()
		var distance = g.mouseEvent.GetDistance(float64(x1), float64(y1))
		fmt.Println(distance)
	}


	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.DrawImage(g.backgroundImage, nil)

	// Draw golf ball
	ebitenutil.DrawCircle(screen, g.ball.x, g.ball.y, g.ball.radius, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Mini Golf")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}