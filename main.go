package main

import (
	"fmt"
	"image/color"
	"log"
	// "math"

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

	// Load golf ball image
	var err error
	var ball_img *ebiten.Image

	ball_img, _, err = ebitenutil.NewImageFromFile("graphics/golf_ball.png")
	if err != nil {
		log.Fatal(err)
	}

	// Create ball struct
	var ball = Ball { 
		x: 250, 
		y: 250, 
		radius: 10, 
		img: ball_img, 
		isMoving: false,
	}

	// Set background
	var backgroundImage = ebiten.NewImage(screenWidth, screenHeight)
	backgroundImage.Fill(color.Black)

	// Create the game struct
	g := &Game { 
		backgroundImage: backgroundImage, 
		ball: ball,
		mouseDown: false,
		mouseEvent: MouseEvent{},
	}

	return g
}

func (g *Game) Update() error {
	// Reset ball position for debugging
	if(inpututil.IsKeyJustPressed(ebiten.KeyR)) {
		fmt.Println("Reset")
		g.ball.x = 250
		g.ball. y = 250
	}

	if(inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)) {
		fmt.Println("Mouse Just Pressed")

		var x0, y0 = ebiten.CursorPosition()
		g.mouseEvent = MouseEvent { x0: float64(x0), y0: float64(y0) }

	} else if (inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)) {
		var x1, y1 = ebiten.CursorPosition()
		g.mouseEvent.SetPower(float64(x1), float64(y1))	
		g.mouseEvent.SetAngle(float64(x1), float64(y1))	

		fmt.Println(g.mouseEvent.angle)

		g.ball.isMoving = true
	}

	if(g.ball.isMoving) {

		if(g.ball.vx0 == 0 && g.ball.vy0 == 0) {
			var angle = g.mouseEvent.angle
			var power = g.mouseEvent.power

			g.ball.SetInitialVelocity(angle, power)
		}

		// fmt.Println("vx: ", velx, " vy: ", vely)

		g.ball.MovePosition()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.DrawImage(g.backgroundImage, nil)

	// Draw golf ball
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.ball.x, g.ball.y)
	screen.DrawImage(g.ball.img, op)
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