package main

import (
	"image/color"
	"log"
	"math"

	"github.com/abolfazlalz/plants_vs_zombie/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game structure to implement the ebiten.Game interface
type Game struct {
	objects []objects.Object
}

// Update is called every frame to update the game state
func (g *Game) Update() error {
	for _, object := range g.objects {
		object.Animate()
	}
	return nil
}

// Draw is called every frame to draw the game on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	plane := objects.NewPlane()
	plane.LoadFrames()

	for i := 0; i < screen.Bounds().Dx(); i += 100 {
		for j := 0; j < screen.Bounds().Dy(); j += 100 {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i), float64(j))
			screen.DrawImage(plane.Frame(i, j), op)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(screen.Bounds().Dy())-100)

	for _, object := range g.objects {
		x, _ := object.Position()
		for _, objectY := range g.objects {
			if objectY == object {
				continue
			}
			x1, _ := objectY.Position()
			// x, x1 = x/100, x1/100
			xfloor := math.Floor(float64(x))
			x1floor := math.Floor(float64(x1))
			if xfloor == x1floor {
				object.Collide(objectY)
			}
		}
		object.Draw(screen)
	}
}

// Layout is called to set the window size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1200, 700
}

func main() {
	ebiten.SetWindowSize(1200, 700)
	ebiten.SetWindowTitle("Game Example with Peashooter Sprite Sheet")

	// Create the Peashooter instance and load its frames
	peashooter := objects.NewPeashooter()
	peashooter.LoadFrames()
	peashooter.SetPosition(1, 5)

	peashooter1 := objects.NewPeashooter()
	peashooter1.LoadFrames()
	peashooter1.SetPosition(2, 2)

	zombie1 := objects.NewZombie()
	zombie1.SetPosition(3, 5)

	zombie2 := objects.NewZombie()
	zombie2.SetPosition(4, 2)

	// Create the game instance and set the Peashooter as part of it
	game := &Game{
		objects: []objects.Object{
			peashooter,
			peashooter1,
			zombie1,
			zombie2,
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
