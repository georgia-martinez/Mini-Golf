package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 500
	screenHeight = 500
)

type Game struct {
	backgroundImage *ebiten.Image
	ball            Ball
	mouseDown       bool
	mouseEvent      MouseEvent
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
	var ball = Ball{
		radius: 10,
		img:    ball_img,
	}
	ball.SetPosition(screenWidth/2, screenHeight/2)

	// Set background
	var backgroundImage = ebiten.NewImage(screenWidth, screenHeight)
	backgroundImage.Fill(GetRGBColor(45, 150, 91))

	// Create the game struct
	g := &Game{
		backgroundImage: backgroundImage,
		ball:            ball,
		mouseEvent:      MouseEvent{},
	}

	return g
}

func (g *Game) Update() error {

	// Defining mouse controls for ball movement
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		var mouseX, mouseY = ebiten.CursorPosition()

		g.mouseEvent = MouseEvent{
			mouseX: float64(mouseX),
			mouseY: float64(mouseY),
		}

		g.ball.canStartMoving = g.ball.IsClicked(g.mouseEvent.mouseX, g.mouseEvent.mouseY) && !g.ball.isMoving
	}

	if g.ball.canStartMoving && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		var x1, y1 = ebiten.CursorPosition()

		g.mouseEvent.SetPower(float64(x1), float64(y1))
		g.mouseEvent.SetAngle(float64(x1), float64(y1))

		// Set the initial velocity
		var angle = g.mouseEvent.angle
		var power = g.mouseEvent.power

		g.ball.SetInitialVelocity(angle, power)
		g.ball.SetInitialDirection()

		g.ball.isMoving = true
	}

	// Update the ball's position
	if g.ball.isMoving {
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

	fmt.Print()

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
