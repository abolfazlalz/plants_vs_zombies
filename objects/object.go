package objects

import (
	"image/png"
	"log"
	"os"
	"sync"

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

type Object interface {
	Animate()
	SetPosition(x, y int)
	Draw(screen *ebiten.Image)
	Collide(Object)
	Position() (x int, y int)
}

type Abstract struct {
	x, y int
	mu   sync.Mutex
}

func NewAbstract() *Abstract {
	return &Abstract{
		x: 0,
		y: 0,
	}
}

func (a *Abstract) Animate() {
	panic("Implement method")
}

func (a *Abstract) SetPosition(x, y int) {
	a.x, a.y = x, y
}

func (a *Abstract) Position() (x, y int) {
	return a.x, a.y
}

func (a *Abstract) Draw(screen *ebiten.Image) {
	panic("Implement method")
}

func (a *Abstract) Collide(obj Object) {
}
