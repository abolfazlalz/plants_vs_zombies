package main

import (
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// Load the sprite sheet from the given path
func loadSpriteSheet(path string) *ebiten.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(file) // Use png.Decode to decode PNG images
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

// Game structure to implement the ebiten.Game interface
type Game struct {
	peashooter *Peashooter
}

// Update is called every frame to update the game state
func (g *Game) Update() error {
	// Handle idle or shot animation based on space key
	if !g.peashooter.isShot {
		g.peashooter.Idle()
	} else {
		g.peashooter.Shot()
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.peashooter.isShot = true
	}

	return nil
}

// Draw is called every frame to draw the game on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.2, 0.2) // Scale image down since each frame is large (500x500)
	screen.DrawImage(g.peashooter.currentFrame, op)
}

// Layout is called to set the window size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(1200, 700)
	ebiten.SetWindowTitle("Game Example with Peashooter Sprite Sheet")

	// Create the Peashooter instance and load its frames
	peashooter := &Peashooter{}
	peashooter.LoadFrames()

	// Create the game instance and set the Peashooter as part of it
	game := &Game{
		peashooter: peashooter,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
